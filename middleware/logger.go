package qrys

import (
	"log"
	"net/http"
	"time"
)

// Log log.printf request info
func Log(w ResponseWriteReader, r *http.Request, next func()) {
	t := time.Now()
	next()
	log.Printf("%v %v %v %v",
		r.Method,
		w.StatusCode(),
		r.URL.String(),
		time.Now().Sub(t).String())
}
