package server

import (
	"net/http"
)

// RouterHandler is the wrapper for NewLogger
type RouterHandler interface {
	GetURLResp(rw http.ResponseWriter, r *http.Request)
	FileScan(rw http.ResponseWriter, r *http.Request)
	MiddlewareValidationForScanRegister(next http.Handler) http.HandlerFunc
	MiddlewareValidationForScan(next http.Handler) http.HandlerFunc
	MiddlewareValidationForFile(next http.Handler) http.HandlerFunc
	MiddlewareForCorsUpdate(next http.Handler) http.HandlerFunc
	PingServer(writer http.ResponseWriter, request *http.Request)
	ScanRegister(writer http.ResponseWriter, request *http.Request)
	GetLatestResults(writer http.ResponseWriter, request *http.Request)
	GetResult(writer http.ResponseWriter, request *http.Request)
	MiddlewareForResults(next http.Handler) http.HandlerFunc
	MiddlewareForResult(next http.Handler) http.HandlerFunc
}
