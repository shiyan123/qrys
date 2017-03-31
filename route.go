package qrys

import (
	"net/http"
)

// Route ...
type Route struct {
	httpMethod string
	locatePath string
	handler    http.Handler
	err        error
}

//Routes Routes []Route
type Routes []Route

// Method ...
func (r *Route) Method(method string) *Route {
	r.httpMethod = method
	return r
}

// LocatePath ...
func (r *Route) LocatePath(locatePath string) *Route {
	r.locatePath = locatePath
	return r
}

// Handler ...
func (r *Route) Handler(handler http.Handler) *Route {
	if r.err == nil {
		r.handler = handler
	}
	return r
}

// HandlerFunc sets a handler function for the route.
func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Handler(http.HandlerFunc(f))
}

// GetHandler ...
func (r *Route) GetHandler() http.Handler {
	return r.handler
}
