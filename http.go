package integration

import "net/http"

// Request wraps a http.Request.
type Request struct {
	*http.Request
}

// ResponseWriter wraps a http.ResponseWriter.
type ResponseWriter struct {
	http.ResponseWriter
}
