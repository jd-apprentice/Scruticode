package config

import (
	"Scruticode/src/constants"
	"Scruticode/src/types"
	"os"
	"strconv"
	"strings"
)

// GetConfig retrieves the configuration for the application from environment variables.
// It returns a Config struct containing Sentry configuration details such as DSN, release,
// environment, and traces sample rate. If the environment variable for the traces sample
// rate is not set or invalid, it defaults to zero.
func GetConfig() types.Config {
	parsedFloat, _ := strconv.ParseFloat(strings.TrimSpace(os.Getenv("SENTRY_TRACES_SAMPLE_RATE")), constants.BitSize)

	return types.Config{
		Sentry: types.SentryConfig{
			TracesSampleRate: parsedFloat,
			Dsn:              os.Getenv("SENTRY_DSN"),
			Release:          os.Getenv("SENTRY_RELEASE"),
			Environment:      os.Getenv("SENTRY_ENVIRONMENT"),
		},
	}
}
