package config

import (
	"fmt"
	"os"
)

type config struct {
	Host      string
	HTTPS     bool
	Port      int
	StripeKey string
}

func (c *config) GetDomain() string {

	protocol := "http"

	if c.HTTPS {
		protocol += "s"
	}

	return fmt.Sprintf("%s://%s:%d", protocol, c.Host, c.Port)
}

func GetConfig() *config {
	return &config{
		Port:      3000,
		HTTPS:     false,
		Host:      "localhost",
		StripeKey: os.Getenv("STRIPE_API_KEY"),
	}
}
