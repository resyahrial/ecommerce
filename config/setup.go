package config

import (
	"fmt"
	"os"
	"time"

	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/migrations"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Config Configuration

func Initialize(env string) {
	setConfig(env)
	initDb()
	initTokenManagerOpts()
}

func Shutdown() {
	shutdownDB()
}

func setConfig(env string) {
	confFilePath := fmt.Sprintf("config/%s.yml", env)
	f, err := os.Open(confFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&Config); err != nil {
		panic(err)
	}

	app.Host = Config.Http.Host
	app.Port = Config.Http.Port

	log.SetOutput(os.Stdout)
}

func initDb() {
	if Config.Db.Host == "" || Config.Db.User == "" {
		panic("warning: no setup for database")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		Config.Db.Host,
		Config.Db.Name,
		Config.Db.Pass,
		Config.Db.Name,
		Config.Db.Port)

	kl := gtrace.NewLogAndTracer(gtrace.LogAndTracer{
		LogLevel:   logger.LogLevel(Config.Db.LogLevel),
		Title:      "sql",
		StringName: "query",
		CountName:  "row",
	})

	var err error
	if app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &kl,
	}); err != nil {
		panic(err)
	}

	sqlDB, err := app.DB.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(Config.Db.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(Config.Db.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(Config.Db.ConnMaxLifetime) * time.Second)

	migrations.AutoMigration(app.DB)
	models.AutoMigrateAllTables(app.DB)
}

func initTokenManagerOpts() {
	timeDurationMap := map[string]time.Duration{
		"hour":   time.Hour,
		"minute": time.Minute,
		"second": time.Second,
	}

	jwtConfig := Config.Jwt
	app.KeyAccess = jwtConfig.KeyAccess
	app.KeyRefresh = jwtConfig.KeyRefresh
	app.ExpiryAgeAccess = time.Duration(jwtConfig.ExpiryAgeAccess.Value) * timeDurationMap[jwtConfig.ExpiryAgeAccess.Unit]
	app.ExpiryAgeRefresh = time.Duration(jwtConfig.ExpiryAgeRefresh.Value) * timeDurationMap[jwtConfig.ExpiryAgeRefresh.Unit]
}

func shutdownDB() {
	log.Info("Closing database connection")
	db, err := app.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
