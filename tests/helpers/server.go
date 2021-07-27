package helpers

import (
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server"

	"github.com/labstack/echo/v4"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/config"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo:   echo.New(),
		DB:     Init(),
		Config: config.NewConfig(),
	}

	return s
}
