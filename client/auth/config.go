package auth

import (
	"errors"
	"os"
)

type Config struct {
	ClientId              string
	ClientSecret          string
	RedirectUri           string
	AuthorizationEndpoint string
	TokenEndpoint         string
	LoginPageEndpoint     string
}

func GetConfig() (*Config, error) {
	config := &Config{}

	clientId := os.Getenv("CLIENT_ID")

	if clientId == "" {
		return nil, errors.New("CLIENT_ID for auth has not been provided")
	}

	config.ClientId = clientId

	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientSecret == "" {
		return nil, errors.New("CLIENT_SECRET for auth has not been provided")
	}

	config.ClientSecret = clientSecret

	redirectUri := os.Getenv("REDIRECT_URI")

	if redirectUri == "" {
		return nil, errors.New("REDIRECT_URI for auth has not been provided")
	}

	config.RedirectUri = redirectUri

	authorizationEndpoint := os.Getenv("AUTHORIZATION_ENDPOINT")

	if authorizationEndpoint == "" {
		return nil, errors.New("AUTHORIZATION_ENDPOINT has not been provided")
	}

	config.AuthorizationEndpoint = authorizationEndpoint

	tokenEndpoint := os.Getenv("TOKEN_ENDPOINT")

	if redirectUri == "" {
		return nil, errors.New("TOKEN_ENDPOINT for auth has not been provided")
	}

	config.TokenEndpoint = tokenEndpoint

	loginPageEndpoint := os.Getenv("LOGIN_PAGE_ENDPOINT")

	if loginPageEndpoint == "" {
		return nil, errors.New("TOKEN_ENDPOINT for auth has not been provided")
	}

	config.LoginPageEndpoint = loginPageEndpoint

	return config, nil
}
