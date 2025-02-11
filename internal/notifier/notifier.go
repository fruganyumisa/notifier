package notifier

import (
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
func (n *Notifier) Start() {
	ticker := time.NewTicker(n.config.CheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		n.CheckServices()
	}
}

// CheckServices checks all services and sends notifications if any are down.
func (n *Notifier) CheckServices() {
	var downServices []string

	for _, service := range n.config.Services {
		n.logger.Info().Str("Event", "Service Monitoring, ").Str("Service", service)

		if !CheckService(service) {
			downServices = append(downServices, service)
		} else {
			n.logger.Info().Str("Event", "Service Monitoring").Str("Status", "Service is up")
		}
	}

	if len(downServices) > 0 {
		message := "Services down: " + joinStrings(downServices, ", ")
		err := SendSMS(n.config.SMSGatewayURL, n.config.AdminPhones, message)
		if err != nil {
			n.logger.Fatal().Str("Event", "Send SMS notification failed ").Msg(err.Error())

		} else {
			log.Println("SMS notification sent.")
			n.logger.Info().Str("Event", "Send SMS notification Successfully").Str("Message", message)
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
