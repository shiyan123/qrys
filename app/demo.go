package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shiyan123/qrys"
	mid "github.com/shiyan123/qrys/middleware"
)

// DemoServer request response
func DemoServer(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}

// Demo1Server request response
func Demo1Server(w http.ResponseWriter, req *http.Request) {
	id := qrys.Vars(req)["id"]
	fmt.Fprintf(w, "RequestURI:%s, method: %s, id: %s", req.URL.RequestURI(), req.Method, id)
}

// Demo2Server request response
func Demo2Server(w http.ResponseWriter, req *http.Request) {
	var jsonData interface{}
	qrys.ParseBody(req, &jsonData)

	fmt.Println(jsonData.(map[string]interface{})["key1"])
	fmt.Println(jsonData.(map[string]interface{})["key2"])
	fmt.Fprintf(w, "RequestURI:%s, method: %s", req.URL.RequestURI(), req.Method)
}

func main() {
	s := new(mid.MiddleWareServe)
	r := qrys.NewRouter()

	r.GET("/", DemoServer)
	r.GET("/a/:id", Demo1Server)
	r.POST("/a", Demo2Server)

	s.Handler = r
	s.Use(mid.Log, mid.ErrCatch)

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
