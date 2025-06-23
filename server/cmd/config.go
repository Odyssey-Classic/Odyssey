package main

import (
	"log"
	"net/url"

	"github.com/caarlos0/env"
)

type Config struct {
	AdminPort   int    `env:"ODY_ADMIN_PORT" envDefault:"8081"`
	MetaPort    int    `env:"ODY_META_PORT" envDefault:"8082"`
	NetworkPort int    `env:"ODY_NETWORK_PORT" envDefault:"3001"`
	ServerID    string `env:"ODY_SERVER_ID,required"`
	APIKey      string `env:"ODY_API_KEY,required"`
	RegistryURL string `env:"ODY_REGISTRY_URL" envDefault:"http://odyssey.local:8080"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// ParseAndValidateURL parses and validates the registry URL.
func ParseAndValidateURL(registry string) *url.URL {
	host, err := url.Parse(registry)
	if err != nil {
		log.Fatal(err)
	}
	if host.Scheme == "" || host.Host == "" {
		log.Fatalf("Invalid registry URL: %s", registry)
	}
	return host
}
