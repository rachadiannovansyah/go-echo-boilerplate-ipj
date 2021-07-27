package main

import (
	"fmt"

	application "github.com/rachadiannovansyah/go-echo-boilerplate-ipj"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/config"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/docs"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
