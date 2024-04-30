package auth

import (
	"encoding/base64"
	"fmt"
	"nekoauth/lib/database"
	"net/http"
	"net/url"
)

func AuthserverUriQuery(r *http.Request, authConfig *Config) (string, *HttpError) {
	state, err := generateState(r, authConfig)

	if err != nil {
		return "", NewHttpError("state generation error: "+err.Error(), http.StatusInternalServerError)
	}

	if err := storeStateInDB(r, state); err != nil {
		return "", NewHttpError("state storage error"+err.Error(), http.StatusInternalServerError)
	}

	redirectQuery := url.Values{}
	redirectQuery.Set("response_type", "code")
	redirectQuery.Set("client_id", authConfig.ClientId) // client identifies itself
	redirectQuery.Set("redirect_uri", authConfig.RedirectUri)
	redirectQuery.Set("state", state)

	// TODO: add scope
	// redirectQuery.Set("scope", authConfig.) // client requests particular items like "scope"

	return redirectQuery.Encode(), nil
}

func HandleRedirectFromAuthserver(r *http.Request, authConfig *Config) *HttpError {
	// parse the query - state, code
	queries := r.URL.Query()

	state := queries.Get("state")
	code := queries.Get("code")

	// compare the state with the one we send to the redirection page at the beginning
	isStateSame, err := compareState(r, state)

	if err != nil {
		NewHttpError(err.Error(), http.StatusInternalServerError)
	}

	if !isStateSame {
		NewHttpError("states are different", http.StatusUnauthorized)
	}

	// asks the authorization server for the token

}

func compareState(r *http.Request, fromAuthserver string) (bool, error) {
	fromDB, err := fetchStateFromDB(r)

	if err != nil {
		return false, err
	}

	return fromDB == fromAuthserver, nil
}

func generateState(r *http.Request, authConfig *Config) (string, error) {
	// TODO: implement logic
	// capture the session-specific data
	// for example, use cookie
	// and encrypt the data with the server-specific data(authConfig)
	return "mock_state", nil
}

func sessionId(r *http.Request) (string, error) {
	// get session id from the given incoming request

	cookie, err := r.Cookie("session_id")

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func storeStateInDB(r *http.Request, state string) error {
	// TODO: use cache database like redis
	db, err := database.DB()

	if err != nil {
		return err
	}

	sessionId, err := sessionId(r)

	if err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO states (id, state) VALUES ($1, $2)`, sessionId, state); err != nil {
		return err
	}

	return nil
}

func fetchStateFromDB(r *http.Request) (string, error) {
	// TODO: use cache db such as redis
	db, err := database.DB()

	if err != nil {
		return "", err
	}

	sessionId, err := sessionId(r)

	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT id, state FROM states WHERE id = ?", sessionId)

	var state string

	if err := row.Scan(&state); err != nil {
		return "", err
	}

	return state, nil
}

func deleteStateInDB(r *http.Request) error {
	// TODO: use cache db such as redis
	db, err := database.DB()

	if err != nil {
		return err
	}

	sessionId, err := sessionId(r)

	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM states WHERE id = ?", sessionId)

	if err != nil {
		return err
	}

	return nil
}

func fetchTokenFromAuthserver(code string, authConfig *Config) error {
	httpClient := http.Client{}

	header := http.Header{}
	header.Set("Accept", "application/json")
	header.Set("Content-type", "application/x-www-form-encoded")

	// base64 encode the following string:
	// <client_id>:<client_secret>

	clientCredential := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf(
			"%s:%s",
			authConfig.ClientId,
			authConfig.ClientSecret,
		)),
	)
	header.Set("Authorization", "Basic "+clientCredential)

	body := url.Values{}
	// the means by which a client is given access to a protected resource using the oauth protocol
	body.Set("grant_type", "authorization_code")
	body.Set("redirect_uri", authConfig.RedirectUri)
	body.Set("code", code)

	url, _ := url.Parse(authConfig.TokenEndpoint)

	clientRequest := http.Request{
		Method:   http.MethodPost,
		URL:      url,
		Header:   header,
		Host:     "/", // client hosts
		PostForm: body,
	}

	resp, err := httpClient.Do(&clientRequest)

	if err != nil || resp.StatusCode >= 400 {
		NewHttpError(err.Error(), http.StatusInternalServerError)
	}

	respBody := resp.Body
	defer respBody.Close()

	// TODO: finish this(3.2.2)
	return nil
}
