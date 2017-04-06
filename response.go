package qrys

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse ...
func NewResponse() *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "OK",
	}
}

// NewResponseWithError ...
func NewResponseWithError(err error) *Response {
	return &Response{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

// Set Message ...
func (resp *Response) Set(msg string) {
	resp.Message = msg
}

// Write json body
func (resp *Response) Write(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.Code)

	body, _ := json.Marshal(resp)
	rw.Write(body)
}
