package main

import (
	"fmt"
	"log"
	"net/http"
)

func DemoServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}
func main() {
	http.HandleFunc("/demo", DemoServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
