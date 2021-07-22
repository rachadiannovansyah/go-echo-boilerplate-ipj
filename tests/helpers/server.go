package helpers

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/server"

	"github.com/khihadysucahyo/go-echo-boilerplate/config"
	"github.com/labstack/echo/v4"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo:   echo.New(),
		DB:     Init(),
		Config: config.NewConfig(),
	}

	return s
}
