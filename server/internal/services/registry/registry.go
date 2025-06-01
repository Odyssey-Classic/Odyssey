package registry

import (
	"log"
	"net/url"
)

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
