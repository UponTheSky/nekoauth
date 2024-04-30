package router

import "net/http"

// This will be moved to a common package in the main backend project
type Router struct {
	method   string
	endpoint string

	handlerFunc http.HandlerFunc
}

func NewRouter(method string, endpoint string, handlerFunc http.HandlerFunc) *Router {
	return &Router{
		method,
		endpoint,
		handlerFunc,
	}
}

func (r *Router) Register(mux *http.ServeMux) {
	mux.HandleFunc(r.method+" "+r.endpoint, r.handlerFunc)
}
