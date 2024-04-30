package server

import (
	"encoding/json"
	"nekoauth/lib/database"
	"net/http"
	"net/url"
	"os"
)

func Run() {
	// bootstrap the test db
	database.BootstrapDB()

	mux := http.NewServeMux()

	redirectUri := "/oauth_callback/"

	// these are temporary: will be made from another server

	mux.HandleFunc("GET /authorize/", func(w http.ResponseWriter, r *http.Request) {
		// step 1: parse the query params

		// step 2: return the template
	})

	mux.HandleFunc("GET /authorize/user/", func(w http.ResponseWriter, r *http.Request) {
		// step 3: from the template, the user identifies oneself to the authorization server
		AuthenticateUser(w, r)
	})

	mux.HandleFunc("GET /authorize/client/", func(w http.ResponseWriter, r *http.Request) {
		// step 4: authorize a scope of authorities to the user
		// the server asks the user whether the one will delegate all the scopes or a part of them
		// the server may cache the info, and might not ask in the future
		AuthorizeClient(w, r)

		// after authorization, the user is redirected to the client application
		oauthAuthCode := "authcode"
		state := "test_state" // the server must hold it when the user is redirected into `/authorize/`
		queries := url.Values{}
		queries.Set("code", oauthAuthCode)
		queries.Set("state", state)
		http.RedirectHandler(redirectUri+"?"+queries.Encode(), http.StatusPermanentRedirect).ServeHTTP(w, r)
	})

	mux.HandleFunc("GET "+redirectUri, func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("POST /authorize/token", func(w http.ResponseWriter, r *http.Request) {
		// step 6: check the credentials and issue a token

		// 6-1: validate the credentials

		// 6-2: read the code params from the body
		// the code contains infos like
		// which client made the init authorization request, which user authorized it, and what it was authorized for

		// if the code is valid, not has been used previously, and the client making this request is the same as
		// the one who made the original request, the auth server generates and returns a new access token for the client

		payload := make(map[string]string)

		// access token includes info like who authorized it and what it was authorized for
		// (usually including the targeted resources)
		// the auth server and the protected resource needs to know how to parse the token
		// whereas the client must be oblivious to it
		payload["access_token"] = "access_token"
		// refresh token is sent to the auth server
		// the clienet can also scope down its access when refreshing the token
		payload["refresh_token"] = "refresh_token"
		payload["token_type"] = "Bearer" // this means that anyone who has this, has the right to use it
		payload["scope"] = "scope"       // a set of rights
		payload["expiry"] = "expiry"

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, mux)
}
