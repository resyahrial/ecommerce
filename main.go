package main

import (
	"flag"

	"github.com/resyahrial/go-commerce/config"
	"github.com/resyahrial/go-commerce/internal/infrastructures/rest"
	api_v1 "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1"
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

	rest.CreateServer(api_v1.Prefix, api_v1.Routes)
}
