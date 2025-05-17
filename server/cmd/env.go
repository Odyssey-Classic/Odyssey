package main

import (
	"os"
	"strconv"
)

// GetUint16 retrieves a uint16 value from the environment or returns the default.
func GetUint16(key string, defaultVal uint16) uint16 {
	if val, ok := os.LookupEnv(key); ok {
		if port, err := strconv.ParseUint(val, 10, 16); err == nil && port <= 65535 {
			return uint16(port)
		}
	}
	return defaultVal
}
