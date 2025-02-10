package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"log"
	"notifier/internal/config"
	"notifier/internal/notifier"
	"os"
)

var (
	c config.Config
)

func main() {
	// Load configuration
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	err := envconfig.Process("NOTIFIER", &c)

	if err != nil {
		logger.Fatal().Str("Event", "Processing service configuration ").Msg(err.Error())
	}

	// Create and start the notifier
	n := notifier.NewNotifier(&c)
	log.Println("Starting notifier service...")
	n.Start()
}
