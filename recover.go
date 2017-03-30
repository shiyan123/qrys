package qrys

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// ErrCatch err
func ErrCatch(w ResponseWriteReader, r *http.Request, next func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	next()
}
