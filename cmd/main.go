package main

import (
	"fmt"

	application "github.com/khihadysucahyo/go-echo-boilerplate"
	"github.com/khihadysucahyo/go-echo-boilerplate/config"
	"github.com/swaggo/swag/example/basic/docs"
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
