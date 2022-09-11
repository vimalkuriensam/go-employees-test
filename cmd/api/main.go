package main

import (
	"flag"

	"github.com/vimalkuriensam/go-employees-test/pkg/config"
)

const DEFAULT_ENVIRONMENT = "development"

var env string

func main() {
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the default environment stage")
	flag.Parse()
	cfg := config.Initialize()
	cfg.LoadEnvironment(env)
	cfg.Logger.Printf("Server is running on port %v", cfg.Env["port"])
}
