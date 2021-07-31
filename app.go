package application

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/config"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server/routes"
)

// Start function
func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.Dsn,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
