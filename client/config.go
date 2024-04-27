package auth

import (
	"errors"
	"os"
)

type Config struct {
	clientId              string
	clientSecret          string
	redirectUri           string
	authorizationEndpoint string
	tokenEndpoint         string
}

func GetConfig() (*Config, error) {
	config := &Config{}

	clientId := os.Getenv("CLIENT_ID")

	if clientId == "" {
		return nil, errors.New("CLIENT_ID for auth has not been provided")
	}

	config.clientId = clientId

	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientSecret == "" {
		return nil, errors.New("CLIENT_SECRET for auth has not been provided")
	}

	config.clientSecret = clientSecret

	redirectUri := os.Getenv("REDIRECT_URI")

	if redirectUri == "" {
		return nil, errors.New("REDIRECT_URI for auth has not been provided")
	}

	config.redirectUri = redirectUri

	authorizationEndpoint := os.Getenv("AUTHORIZATION_ENDPOINT")

	if authorizationEndpoint == "" {
		return nil, errors.New("AUTHORIZATION_ENDPOINT has not been provided")
	}

	config.authorizationEndpoint = authorizationEndpoint

	tokenEndpoint := os.Getenv("TOKEN_ENDPOINT")

	if redirectUri == "" {
		return nil, errors.New("TOKEN_ENDPOINT for auth has not been provided")
	}

	config.tokenEndpoint = tokenEndpoint

	return config, nil
}
