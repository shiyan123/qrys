package qrys

import (
	"net/http"
)

// NewRouter ...
func NewRouter() *Router {
	return &Router{make(map[string][]*Handler), nil}
}

// Router ...
type Router struct {
	handlers map[string][]*Handler
	NotFound http.HandlerFunc
}

// Handler ...
type Handler struct {
	pattern string
	http.HandlerFunc
}

// Listen ...
func (r *Router) Listen(port string) {
	http.ListenAndServe(port, r)
}

//  implement ServeHTTP
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pathLen := len(req.URL.Path)

	// Redirect ...req.URL.Path == "/"
	if pathLen > 1 && req.URL.Path[pathLen-1:] == "/" {
		http.Redirect(w, req, req.URL.Path[:pathLen-1], 301)
		return
	}

	for _, handler := range r.handlers[req.Method] {
		if handler.pattern == req.URL.Path {
			handler.ServeHTTP(w, req)
			return
		}
	}
	if r.NotFound != nil {
		w.WriteHeader(404)
		r.NotFound.ServeHTTP(w, req)
		return
	}
	// 404
	http.NotFound(w, req)
}

//add route ...
func (r *Router) add(method, pattern string, handler http.HandlerFunc) {
	h := &Handler{
		pattern,
		handler,
	}

	r.handlers[method] = append(
		r.handlers[method],
		h,
	)
}

// GET adds a new route for GET requests.
func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.add("GET", pattern, handler)
}

// POST adds a new route for POST requests.
func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.add("POST", pattern, handler)
}
