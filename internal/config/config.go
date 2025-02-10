package config

import (
	"github.com/rs/zerolog"
	"time"
)

type Config struct {
	Services      []string      `required:"true" split_words:"true"`
	CheckInterval time.Duration `required:"true" split_words:"true"`
	SMSGatewayURL string        `required:"true" split_words:"true"`
	AdminPhones   []string      `required:"true" split_words:"true"`
	Logger        zerolog.Logger
}
