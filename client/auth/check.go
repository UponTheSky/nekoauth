package auth

import (
	"net/http"
)

func AuthenticateUser(r *http.Request, authConfig *Config) (bool, *HttpError) {
	// currently we use cookie for token persistance
	cookie, err := r.Cookie("token")

	if err != nil {
		return false, NewHttpError(err.Error(), http.StatusBadRequest)
	}

	isAuthenticated, err := checkAuthserver(cookie.Value, authConfig)

	if err != nil {
		return false, NewHttpError(err.Error(), http.StatusInternalServerError)
	}

	return isAuthenticated, nil
}

func checkAuthserver(token string, authConfig *Config) (bool, error) {

	// TODO: implement this
	return true, nil
}
