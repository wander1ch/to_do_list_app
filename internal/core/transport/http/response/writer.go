package core_http_response

import (
	"net/http"
)

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		StatusCode: StatusCodeUninitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriter) GetStatusCode() int {
	if rw.StatusCode == StatusCodeUninitialized {
		panic("StatusCode not set. WriteHeader must be called before GetStatusCode.") 
	}
	return rw.StatusCode
}
