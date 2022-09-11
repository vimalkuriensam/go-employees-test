package config

import (
	"log"
	"time"
)

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
