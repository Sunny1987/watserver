package server

import (
	"net/http"
)

// RouterHandler is the wrapper for NewLogger
type RouterHandler interface {
	GetURLResp(rw http.ResponseWriter, r *http.Request)
	FileScan(rw http.ResponseWriter, r *http.Request)
	MiddlewareValidationForURL(next http.Handler) http.HandlerFunc
	MiddlewareValidationForFile(next http.Handler) http.HandlerFunc
	MiddlewareForCorsUpdate(next http.Handler) http.HandlerFunc
}
