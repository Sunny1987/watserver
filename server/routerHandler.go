package server

import (
	"net/http"
)

// RouterHandler is the wrapper for NewLogger
type RouterHandler interface {
	GetURLResp(rw http.ResponseWriter, r *http.Request)
	FileScan(rw http.ResponseWriter, r *http.Request)
	MiddlewareValidation(next http.Handler) http.Handler
}
