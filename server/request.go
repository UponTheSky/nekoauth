package server

import "net/http"

type OAuth2RequestHeader struct {
	ResponseType string
	Scope        string
	ClientId     string
	RedirectUri  string
	State        string
}

func ParseHeader(r *http.Request) OAuth2RequestHeader {
	responseType := r.Header.Get("response_type")
	scope := r.Header.Get("scope")
	clientId := r.Header.Get("client_id")
	redirectUri := r.Header.Get("redirect_uri")
	state := r.Header.Get("state")

	return OAuth2RequestHeader{
		ResponseType: responseType,
		Scope:        scope,
		ClientId:     clientId,
		RedirectUri:  redirectUri,
		State:        state,
	}
}

func AuthorizeClient(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized client", http.StatusUnauthorized)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized user", http.StatusUnauthorized)
}

func GrantClient(w http.ResponseWriter, r *http.Request) {
	// TODO: cache the user config in the DB and skip the grant process
	http.Redirect(w, r, "REDIRECT_URI", http.StatusFound)
}
