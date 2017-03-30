package main

import (
	"fmt"
	"log"
	"net/http"
	"qrys"
)

// DemoServer request response
func DemoServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}

func main() {
	s := new(qrys.MiddleWareServe)
	route := http.NewServeMux()

	route.Handle("/sy", http.HandlerFunc(DemoServer))

	s.Handler = route
	s.Use(qrys.Log, qrys.ErrCatch)

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
