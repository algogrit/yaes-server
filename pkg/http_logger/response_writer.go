package httplogger

import "net/http"

// ResponseWriter is a wrapper over http.RResponseWriter to record the statusCode
type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader is a wrapper for http.ResponseWriter's WriteHeader method
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// NewResponseWriter returns an instance of ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &ResponseWriter{w, http.StatusOK}
}
