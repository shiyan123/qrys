package qrys

import (
	"encoding/json"
	"io/ioutil"
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
	pattern  string
	patterns []string
	wild     bool
	http.HandlerFunc
}

var vars = map[*http.Request]map[string]string{}

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
	//split URL
	segments := split(trim(req.URL.Path, "/"), "/")
	for _, handler := range r.handlers[req.Method] {
		if handler.pattern == req.URL.Path {
			handler.ServeHTTP(w, req)
			return
		}
		if ok, v := handler.Try(segments); ok {
			vars[req] = v
			handler.ServeHTTP(w, req)
			delete(vars, req)
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
		split(trim(pattern, "/"), "/"),
		pattern[len(pattern)-1:] == "*",
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

// Vars GET URL Query
func Vars(r *http.Request) map[string]string {
	if v, ok := vars[r]; ok {
		return v
	}
	return nil
}

// ParseBody ...
func ParseBody(req *http.Request, container *interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &container)
	if err != nil {
		return err
	}
	return nil
}

// JSONBody ...
func JSONBody(data interface{}, code int, rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	body, err := json.Marshal(data)
	if err != nil {
		body, _ = json.Marshal(NewResponseWithError(err))
	}
	rw.Write(body)
}

// Try ...
func (h *Handler) Try(usegs []string) (bool, map[string]string) {
	if len(h.patterns) != len(usegs) && !h.wild {
		return false, nil
	}

	vars := map[string]string{}

	for idx, part := range h.patterns {
		if part == "*" {
			continue
		}
		if part != "" && part[0] == ':' {
			vars[part[1:]] = usegs[idx]
			continue
		}
		if part != usegs[idx] {
			return false, nil
		}
	}

	return true, vars
}

// trim ...
func trim(s, w string) string {
	ss := len(s)
	ws := len(w)
	if ss >= ws && s[:ws] == w {
		s = s[ws:]
		ss -= ws
	}
	if ss >= ws && s[ss-ws:] == w {
		s = s[:ss-ws]
	}
	return s
}

// split ...
func split(s, d string) []string {
	var i, c, n int
	l := len(s)
	for i < l {
		if s[i:i+1] == d {
			c++
		}
		i++
	}
	a := make([]string, c+1)
	c, i = 0, 0
	for i < l {
		if i == l-1 {
			a[c] = s[n : i+1]
			break
		}
		if s[i:i+1] == d {
			a[c] = s[n:i]
			n = i + 1
			c++
		}
		i++
	}
	return a
}
