package config

import (
	"Scruticode/src/constants"
	"Scruticode/src/types"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// GetConfig retrieves the configuration for the application from environment variables.
// It returns a Config struct containing Sentry configuration details such as DSN, release,
// environment, and traces sample rate. If the environment variable for the traces sample
// rate is not set or invalid, it defaults to zero.
func GetConfig() (types.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sampleRateStr := os.Getenv("SENTRY_TRACES_SAMPLE_RATE")

	if sampleRateStr == "" {
		sampleRateStr = "1.0"
	}

	parsedFloat, err := strconv.ParseFloat(strings.TrimSpace(sampleRateStr), constants.BitSize)

	if err != nil {
		return types.Config{}, err
	}

	return types.Config{
		Sentry: types.SentryConfig{
			TracesSampleRate: parsedFloat,
			Dsn:              os.Getenv("SENTRY_DSN"),
			Release:          os.Getenv("SENTRY_RELEASE"),
			Environment:      os.Getenv("SENTRY_ENVIRONMENT"),
		},
	}, nil
}
