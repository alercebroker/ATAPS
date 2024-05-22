package tapsync

import "os"

// Config is the configuration for the application
type Config struct {
	DatabaseURL string
}

func GetConfig() *Config {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
