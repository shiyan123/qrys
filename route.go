package qrys

import (
	"net/http"
)

// Handler ...
type Handler func(*http.Request) *http.ResponseWriter

//Route ...
type Route struct {
	httpMethod string
	locatePath string
	handler    Handler
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

// func (routes *Routes) RegisterUrl(method string, uri string, h http.HandleFunc) *Routes {
// 	route := Route{}
// 	route.LocatePath(uri).Method(method)
// 	route.handler = h

// 	*routes = append(*routes, route)
// 	return routes
// }
