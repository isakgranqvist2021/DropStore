package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host      string
	HTTPS     bool
	Port      int
	StripeKey string
}

func (c *Config) GetDomain() string {

	protocol := "http"

	if c.HTTPS {
		protocol += "s"
	}

	return fmt.Sprintf("%s://%s:%d", protocol, c.Host, c.Port)
}

func GetConfig() *Config {
	return &Config{
		Port:      3000,
		HTTPS:     false,
		Host:      "localhost",
		StripeKey: os.Getenv("STRIPE_API_KEY"),
	}
}
