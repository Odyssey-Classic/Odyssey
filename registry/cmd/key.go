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
)

func getKeyPath() string {
	if envPath := os.Getenv("ODY_PRIVATE_KEY_PATH"); envPath != "" {
		return envPath
	}
	home, err := os.UserHomeDir()
	if err != nil {
		home = "." // fallback to current dir
	}
	return home + "/.odyssey_registry_private_key.pem"
}

func loadOrCreateECDSAKey() (*ecdsa.PrivateKey, error) {
	keyFile := getKeyPath()
	if _, err := os.Stat(keyFile); err == nil {
		pemBytes, err := os.ReadFile(keyFile)
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
	slog.Info("[registry] No private key found, generating new ECDSA key", "path", keyFile)
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
	if err := os.WriteFile(keyFile, pemBytes, 0600); err != nil {
		return nil, err
	}
	return key, nil
}
