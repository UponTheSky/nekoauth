package router

import (
	"nekoauth/client/auth"
	"net/http"
)

func RegisterRouters(mux *http.ServeMux, authConfig *auth.Config) {
	routers := []*Router{}

	// show the login page to the user
	routers = append(routers, NewRouter(
		http.MethodPost,
		authConfig.LoginPageEndpoint,
		func(w http.ResponseWriter, r *http.Request) {
			// TODO: return template
		}))

	// redirect the user to the authserver
	routers = append(routers, NewRouter(
		http.MethodGet,
		"/auth/login/authserver/",
		func(w http.ResponseWriter, r *http.Request) {
			authserverUriQuery, err := auth.AuthserverUriQuery(r, authConfig)

			if err != nil {
				http.Error(w, err.Error(), err.Status)
			}

			endpoint := authConfig.AuthorizationEndpoint + "?" + authserverUriQuery
			http.Redirect(w, r, endpoint, http.StatusFound)
		}))

	// when the user is redirected from the authserver, parse code
	// and get a token from the authserver
	// and finally redirect again to the main page
	routers = append(routers, NewRouter(
		http.MethodGet,
		authConfig.RedirectUri,
		func(w http.ResponseWriter, r *http.Request) {

			if err := auth.HandleRedirectFromAuthserver(r, authConfig); err != nil {
				http.Error(w, err.Error(), err.Status)
			}

			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		}))

	for _, router := range routers {
		router.Register(mux)
	}
}
