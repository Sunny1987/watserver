package server

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
	"strings"
	"time"
)

// MiddlewareValidationForScanRegister will perform validations for ScanRegister
func (l *NewLogger) MiddlewareValidationForScanRegister(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareValidationForScanRegister***")
		if request.URL.Path == "/api/v1/scanregister" {
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
			ctx, cancel := context.WithTimeout(request.Context(), 5*time.Second)
			defer cancel()
			request = request.WithContext(ctx)

			request = request.WithContext(context.WithValue(request.Context(), "req", req))

		}
		l.myLogger.Println("*****Exiting MiddlewareValidationForScanRegister******")
		next.ServeHTTP(writer, request)
	}
}

// MiddlewareValidationForScan will perform validations for GetURLResp
func (l *NewLogger) MiddlewareValidationForScan(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareValidationForScan***")
		if request.URL.Path == "/api/v1/scan" {
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
			ctx, cancel := context.WithTimeout(request.Context(), 5*time.Second)
			defer cancel()
			request = request.WithContext(ctx)

			request = request.WithContext(context.WithValue(request.Context(), "fReq", req))

		}
		l.myLogger.Println("*****Exiting MiddlewareValidationForScan******")
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
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Control-Allow-Origin")
		writer.WriteHeader(http.StatusOK)

		l.myLogger.Println("***Exiting MiddlewareForCorsUpdate***")
		next.ServeHTTP(writer, request)
	}
}

func (l *NewLogger) MiddlewareForResults(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareForResults***")
		if request.URL.Path == "/api/v1/results" {
			if request.Method != http.MethodGet {
				http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}

		next.ServeHTTP(writer, request)
	}
}

func (l *NewLogger) MiddlewareForResult(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting MiddlewareForResult***")

		if strings.Contains(request.URL.Path, "/api/v1/result") {
			if id := strings.Split(request.URL.Path, "/")[2]; id == "" {
				http.Error(writer, "Path does not contain id", http.StatusBadRequest)
				return
			}
		}
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
