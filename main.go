package main

import (
	"flag"

	"github.com/resyahrial/go-commerce/config"
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
}
