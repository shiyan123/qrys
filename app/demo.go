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

// Demo1Server request response
func Demo1Server(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}

// Demo2Server request response
func Demo2Server(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}

func main() {
	s := new(qrys.MiddleWareServe)
	r := qrys.NewRouter()

	r.GET("/", DemoServer)
	r.GET("/a", Demo1Server)
	r.GET("/a/a", Demo1Server)
	r.POST("/a", Demo2Server)

	s.Handler = r
	s.Use(qrys.Log, qrys.ErrCatch)

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
