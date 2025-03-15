package main

import (
	"Scruticode/src/config"
	"Scruticode/src/constants"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	config.InitSentry()
	println("Hello, World!")

	defer sentry.Flush(constants.FlushTime * time.Second)
	sentry.CaptureMessage("Scruticode started")
}
