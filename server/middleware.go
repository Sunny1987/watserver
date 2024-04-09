package server

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
	"strings"
)

// MiddlewareValidationForURL will perform validations for GetURLResp
func (l *NewLogger) MiddlewareValidationForURL(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareForURL***")
		if request.URL.Path == "/api/v1/scan" || request.URL.Path == "/api/v1/scanregister" {
			req := &MyRequest{}
			err := json.NewDecoder(request.Body).Decode(req)
			if err != nil {
				l.myLogger.Println("Middleware: %v", err)
			}
			err = req.Validate()
			if err != nil {
				if strings.Contains(err.Error(), "lte") {
					http.Error(writer, "Depth is greater than 3", http.StatusBadRequest)
				}
				if strings.Contains(err.Error(), "gt") {
					http.Error(writer, "Depth is less than 0", http.StatusBadRequest)
				}

			}
		}
		l.myLogger.Println("*****Exiting MiddlewareForURL******")
		next.ServeHTTP(writer, request)
	}
}

// MiddlewareValidationForFile will perform validations for FileScan
func (l *NewLogger) MiddlewareValidationForFile(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareValidationForFile***")
		if request.URL.Path == "/uploadhtml" {
			err := request.ParseMultipartForm(10 << 20)
			if err != nil {
				http.Error(writer, "Error max file size exceeded", http.StatusBadRequest)
				return
			}

		}
		l.myLogger.Println("***Exiting MiddlewareValidationForFile***")
		next.ServeHTTP(writer, request)
	}
}

// MiddlewareForCorsUpdate will provide CORS access
func (l *NewLogger) MiddlewareForCorsUpdate(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareForCorsUpdate***")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		l.myLogger.Println("***Exiting MiddlewareForCorsUpdate***")
		next.ServeHTTP(writer, request)
	}
}

// ValidateRequestURL will validate url object empty and http/https protocols
func ValidateRequestURL(fl validator.FieldLevel) bool {

	//check if url is empty
	if fl.Field().String() == "" {
		return false
	}

	//check if url has http/https
	if !strings.Contains(fl.Field().String(), "https://") &&
		!strings.Contains(fl.Field().String(), "http://") {
		return false
	}
	return true
}
