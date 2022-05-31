package migrations

import (
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/hasher"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

func seedUser(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&models.User{}) {
		return
	}

	if err = db.AutoMigrate(
		models.User{},
	); err != nil {
		return
	}

	hashHelper := hasher.NewBcyptHasher()
	buyerPassword, _ := hashHelper.Hash("buyer123")
	sellerPassword, _ := hashHelper.Hash("seller123")

	return db.Create(&[]models.User{
		{
			ID:       ksuid.New(),
			Email:    "buyer@gmail.com",
			Name:     "Buyer",
			Password: buyerPassword,
			Address:  "Jl. Bersama, Kota Malang, Jawa Timur",
			Role:     "BUYER",
		},
		{
			ID:       ksuid.New(),
			Email:    "seller@gmail.com",
			Name:     "Seller",
			Password: sellerPassword,
			Address:  "Jl. Sendiri, Kota Bekasi, Jawa Barat",
			Role:     "SELLER",
		},
	}).Error
}
