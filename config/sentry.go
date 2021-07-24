package config

import "os"

type SentryConfig struct {
	Dsn              string
	TracesSampleRate float64
}

func LoadSentryConfig() SentryConfig {
	return SentryConfig{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	}
}
