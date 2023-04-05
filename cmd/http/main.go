package main

import (
	"os"
	"time"

	log "github.com/charmbracelet/log"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	logger.Info("Hello World!")
}
