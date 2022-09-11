package config

import (
	"log"
	"os"
	"time"
)

var cfg *Config

type Config struct {
	DataChan      chan any
	Logger        *log.Logger
	Response      *JSONResponse
	ErrorResponse *ErrorResponse
}

type JSONResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Status    int       `json:"status"`
	Path      string    `json:"path"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}

func Initialize() *Config {
	cfg = &Config{
		DataChan:      make(chan any),
		Logger:        log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Response:      &JSONResponse{},
		ErrorResponse: &ErrorResponse{},
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}
