package server

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
	"strings"
)

type KeyUser struct{}

var ctx context.Context

// MiddlewareValidation will validate incoming request and authentication before scan
func (l *NewLogger) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		l.myLogger.Println("***Starting Middleware***")
		req := &MyRequest{}
		err := json.NewDecoder(request.Body).Decode(req)
		if err != nil {
			l.myLogger.Println("Middleware: %v", err)
		}
		switch request.URL.Path {
		case "/scan":
			err = req.Validate()
			if err != nil {
				if strings.Contains(err.Error(), "lte") {
					http.Error(writer, "Depth is greater than 3", http.StatusBadRequest)
				}
				if strings.Contains(err.Error(), "gt") {
					http.Error(writer, "Depth is less than 0", http.StatusBadRequest)
				}
				l.myLogger.Println("*****Exiting middleware******")
			}
			ctx = context.WithValue(request.Context(), KeyUser{}, req)
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)

		case "/uploadhtml":
			ctx = context.WithValue(request.Context(), KeyUser{}, req)
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		}
	})
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
