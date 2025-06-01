package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log/slog"
	"os"
	"path"
)

func defaultKeyPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "." // fallback to current dir
	}
	path := path.Join(home, ".odyssey_registry_private_key.pem")
	return path
}

// loadOrCreateECDSAKey loads an ECDSA private key from the given path, or generates and saves a new one if not found.
func loadOrCreateECDSAKey(path string) (*ecdsa.PrivateKey, error) {
	if _, err := os.Stat(path); err == nil {
		pemBytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		block, _ := pem.Decode(pemBytes)
		if block == nil || block.Type != "EC PRIVATE KEY" {
			return nil, errors.New("failed to decode PEM block containing EC private key")
		}
		key, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return key, nil
	}
	// Key does not exist, generate and save
	slog.Info("[registry] No private key found, generating new ECDSA key", "path", path)
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}
	pemBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}
	pemBytes := pem.EncodeToMemory(pemBlock)
	if err := os.WriteFile(path, pemBytes, 0600); err != nil {
		return nil, err
	}
	return key, nil
}
