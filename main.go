package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/resyahrial/go-commerce/config"
	"github.com/resyahrial/go-commerce/internal/infrastructures/rest"
	"github.com/sirupsen/logrus"
)

func main() {
	var env string
	flag.StringVar(&env,
		"env",
		"example",
		"env of deployment, will load the respective yml conf file.",
	)
	flag.Parse()

	config.Initialize(env)
	defer config.Shutdown()

	go func() {
		rest.CreateServer()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		logrus.Warn(sig)
		done <- true
	}()
	<-done
}
