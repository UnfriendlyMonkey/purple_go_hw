package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
