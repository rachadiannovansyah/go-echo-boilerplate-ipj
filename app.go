package application

import (
	"log"

	"github.com/khihadysucahyo/go-echo-boilerplate/config"
	"github.com/khihadysucahyo/go-echo-boilerplate/server"
	"github.com/khihadysucahyo/go-echo-boilerplate/server/routes"
)

// Start func
func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
