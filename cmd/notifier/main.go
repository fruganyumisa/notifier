package main

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"log"
	"notifier/internal/config"
	"notifier/internal/notifier"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// Create a context that cancels on termination signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Create and start the notifier
	n := notifier.NewNotifier(&c)
	go func() {
		logger.Info().Str("Start Service", "Notifier Service starting").Msg("Check Service called...")
		n.Start(ctx)
	}()

	// Wait for a termination signal
	sig := <-signalChan
	log.Printf("Received signal: %v. Shutting down...\n", sig)

	// Cancel the context to stop the notifier
	cancel()

	// Allow some time for cleanup
	time.Sleep(1 * time.Second)
	logger.Info().Str("Shutdown Service", "Notifier Service shutdown").Msg("Check Service shutting down")

}
