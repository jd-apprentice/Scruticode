package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {
	config, err := GetConfig()

	if err != nil {
		log.Fatalf("GetConfig: %s", err)
	}

	if config.Sentry.Environment == "prod" {
		sentry.Init(sentry.ClientOptions{
			Dsn:              config.Sentry.Dsn,
			TracesSampleRate: config.Sentry.TracesSampleRate,
			Release:          config.Sentry.Release,
			Environment:      config.Sentry.Environment,
		})
	}
}
