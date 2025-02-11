package notifier

import (
	"context"
	"github.com/rs/zerolog"
	"log"
	"notifier/internal/config"
	"os"
	"time"
)

// Notifier monitors services and sends notifications.
type Notifier struct {
	config *config.Config
	logger zerolog.Logger
}

// NewNotifier creates a new Notifier instance.
func NewNotifier(cfg *config.Config) *Notifier {
	return &Notifier{config: cfg, logger: zerolog.New(os.Stdout).With().Timestamp().Logger()}
}

// Start begins the monitoring process.
func (n *Notifier) Start(ctx context.Context) {
	ticker := time.NewTicker(n.config.CheckInterval)
	n.logger.Info().Str("Service Start", "Notifier Started").Msg("Waiting for the first tick")
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Tick received. Checking services...")
			n.logger.Info().Str("Tick status", "Ticker Received").Msg("Tick received attempting Service check")
			n.CheckServices()
		case <-ctx.Done():
			n.logger.Info().Str("Graceful Shutdown", "Received Shutdown Signal").Msg("Notifier shutting down in few seconds...")
			return
		}
	}

}

// CheckServices checks all services and sends notifications if any are down.
func (n *Notifier) CheckServices() {
	var downServices []string

	for _, service := range n.config.Services {
		n.logger.Info().Str("Event", "Service Monitoring, ").Str("Service", service).Msg("Service check completed")

		if !CheckService(service) {
			downServices = append(downServices, service)
		} else {
			n.logger.Info().Str("Event", "Service Monitoring").Str("Status", "Service is up").Msg(service)
		}
	}

	if len(downServices) > 0 {
		message := "Hello Admin, \nCritical services detected to be down details of Services down: \n" + joinStrings(downServices, ", \n")
		err := SendSMS(n.config.SMSGatewayURL, n.config.AdminPhones, message, n.config.SenderHeader)
		if err != nil {
			n.logger.Fatal().Str("Event", "Send SMS notification failed ").Msg(err.Error())

		} else {
			log.Println("SMS notification sent.")
			n.logger.Info().Str("Event", "Send SMS ").Str("Message", message).Msg("SMS notification sent Successfully")
		}
	}
}

// joinStrings joins a slice of strings with a delimiter.
func joinStrings(strings []string, delimiter string) string {
	result := ""
	for i, s := range strings {
		if i > 0 {
			result += delimiter
		}
		result += s
	}
	return result
}
