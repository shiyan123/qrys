package qrys

import (
	"net/http"
)

// NewRouter ...
func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route), server: &http.Server{}}
}

// Router ...
type Router struct {
	server      *http.Server
	routes      []*Route
	namedRoutes map[string]*Route
}

// import ServeHTTP
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// func (router *Router) Run(port string) error {
// 	router.server.Handler = router.routes
// 	router.server.Addr = port
// 	return router.server.ListenAndServe()
// }
