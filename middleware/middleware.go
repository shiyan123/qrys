package qrys

import (
	"net/http"
)

// MiddleWareFunc type
type MiddleWareFunc func(ResponseWriteReader, *http.Request, func())

// MiddleWareServe type
type MiddleWareServe struct {
	middleWares []MiddleWareFunc
	Handler     http.Handler
}

// ResponseWriteReader : 1.code 2.http.ResponseWriter
type ResponseWriteReader interface {
	StatusCode() int
	http.ResponseWriter
}

// WrapResponseWriter implement ResponseWriteReader interface
type WrapResponseWriter struct {
	status int
	http.ResponseWriter
}

// NewWrapResponseWriter create wrapResponseWriter
func NewWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	wr := new(WrapResponseWriter)
	wr.ResponseWriter = w
	wr.status = 200
	return wr
}

// WriteHeader write status code
func (p *WrapResponseWriter) WriteHeader(status int) {
	p.status = status
	p.ResponseWriter.WriteHeader(status)
}

func (p *WrapResponseWriter) Write(b []byte) (int, error) {
	n, err := p.ResponseWriter.Write(b)
	return n, err
}

// StatusCode return status code
func (p *WrapResponseWriter) StatusCode() int {
	return p.status
}

//implement ServeHTTP
func (m *MiddleWareServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := 0
	// warp http.ResponseWriter 可以让中间件读取到 status code
	wr := NewWrapResponseWriter(w)
	var next func() // next 函数指针
	next = func() {
		if i < len(m.middleWares) {
			i++
			m.middleWares[i-1](wr, r, next)
		} else if m.Handler != nil {
			m.Handler.ServeHTTP(wr, r)
		}
	}
	next()
}

// Use insert MiddleWareFunc 可以插入多个MiddleWareFunc
func (m *MiddleWareServe) Use(funcs ...MiddleWareFunc) {
	for _, f := range funcs {
		m.middleWares = append(m.middleWares, f)
	}
}
