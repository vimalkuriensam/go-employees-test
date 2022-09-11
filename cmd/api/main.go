package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/go-employees-test/pkg/config"
	"github.com/vimalkuriensam/go-employees-test/pkg/routes"
)

const DEFAULT_ENVIRONMENT = "development"

var env string

func main() {
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the default environment stage")
	flag.Parse()
	cfg := config.Initialize()
	cfg.LoadEnvironment(env)
	err := cfg.MongoConnect()
	if err != nil {
		cfg.Logger.Panic(err)
	}
	defer cfg.MongoDisconnect()
	cfg.InsertMongoCollections("employees")
	routes := routes.Routes()
	cfg.Logger.Printf("Server is running on port %v", cfg.Env["port"])
	cfg.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Env["port"]), routes))
}
