package tapsync

import "os"

// Config is the configuration for the application
type Config struct {
	DatabaseURL string
	Port        int
}

type ConfigOption func(*Config)

func NewConfig(opts ...ConfigOption) *Config {
	defaultDatabaseUrl := os.Getenv("DATABASE_URL")
	defaultPort := 8080
	config := &Config{
		DatabaseURL: defaultDatabaseUrl,
		Port:        defaultPort,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func WithDatabaseURL(databaseUrl string) ConfigOption {
	return func(c *Config) {
		c.DatabaseURL = databaseUrl
	}
}

func WithPort(port int) ConfigOption {
	return func(c *Config) {
		c.Port = port
	}
}
