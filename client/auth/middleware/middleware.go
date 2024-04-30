package middleware

import (
	"nekoauth/client/auth"
	"net/http"
)

func AuthUserMiddleware(mux http.Handler, authConfig *auth.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isAuthenticated, httpErr := auth.AuthenticateUser(r, authConfig)

		if httpErr != nil {
			http.Error(w, httpErr.Error(), httpErr.Status)
		}

		if !isAuthenticated {
			http.Redirect(w, r, authConfig.LoginPageEndpoint, http.StatusPermanentRedirect)
		}

		mux.ServeHTTP(w, r)
	})
}
