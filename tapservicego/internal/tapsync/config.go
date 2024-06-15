package tapsync

import "os"

// Config is the configuration for the application
type Config struct {
	DatabaseURL string
}

func GetConfig() *Config {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("DATABASE_URL not set")
	}
	return &Config{
		DatabaseURL: databaseUrl,
	}
}
