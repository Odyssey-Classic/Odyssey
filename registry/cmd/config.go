package main

import (
	"errors"
	"net/url"
	"os"
)

const (
	authVar         = "ODY_AUTHORIZATION_URL"
	tokenVar        = "ODY_TOKEN_URL"
	revokeVar       = "ODY_REVOKE_URL"
	redirectVar     = "ODY_REDIRECT_URL"
	clientIDVar     = "ODY_CLIENT_ID"
	clientSecretVar = "ODY_CLIENT_SECRET"
)

type OAuthConfig struct {
	AuthorizationURL *url.URL
	RedirectURL      *url.URL
	RevokeURL        *url.URL
	TokenURL         *url.URL

	ClientID     string
	ClientSecret string
}

type Config struct {
	OAuth OAuthConfig
}

func ConfigFromEnv() (*Config, error) {
	var allErr error
	authEnv, ok := os.LookupEnv(authVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_AUTHORIZATION_URL is required"))
	}
	authURL, err := url.Parse(authEnv)
	if err != nil {
		allErr = errors.Join(err, errors.New("ODY_AUTHORIZATION_URL is invalid"))
	}

	redirectEnv, ok := os.LookupEnv(redirectVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_REDIRECT_URL is required"))
	}
	redirectURL, err := url.Parse(redirectEnv)
	if err != nil {
		allErr = errors.Join(allErr, err)
	}

	revokeEnv, ok := os.LookupEnv(revokeVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_REVOKE_URL is required"))
	}
	revokeURL, err := url.Parse(revokeEnv)
	if err != nil {
		allErr = errors.Join(allErr, err)
	}

	tokenEnv, ok := os.LookupEnv(tokenVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_TOKEN_URL is required"))
	}
	tokenURL, err := url.Parse(tokenEnv)
	if err != nil {
		allErr = errors.Join(allErr, err)
	}

	clientID, ok := os.LookupEnv(clientIDVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_CLIENT_ID is required"))
	}

	clientSecret, ok := os.LookupEnv(clientSecretVar)
	if !ok {
		allErr = errors.Join(allErr, errors.New("ODY_CLIENT_SECRET is required"))
	}

	if allErr != nil {
		return nil, allErr
	}

	cfg := &Config{
		OAuth: OAuthConfig{
			AuthorizationURL: authURL,
			TokenURL:         tokenURL,
			RedirectURL:      redirectURL,
			RevokeURL:        revokeURL,
			ClientID:         clientID,
			ClientSecret:     clientSecret,
		},
	}

	return cfg, nil
}
