package server

import (
	"github.com/go-playground/validator"
	"io"
	"log"
	"net/http"
	"webserver/parserapp/parser"
	"webserver/resultsapp"
)

// MyRequest object for input
type MyRequest struct {
	URL          string               `json:"url" validate:"required,url"`
	Depth        int                  `json:"depth" validate:"gt=-1,lte=2"`
	File         io.Reader            `json:"file"`
	FileName     string               `json:"fileName"`
	parserEngine parser.ParserService `json:"-"`
}

func NewMyRequestFile(file io.Reader, fileName string, parserEngine parser.ParserService) *MyRequest {
	return &MyRequest{File: file, FileName: fileName, parserEngine: parserEngine}
}

func NewMyRequestURL(URL string, depth int, parserEngine parser.ParserService) *MyRequest {
	return &MyRequest{URL: URL, Depth: depth, parserEngine: parserEngine}
}

// Validate will perform field level validation
func (req *MyRequest) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("url", ValidateRequestURL)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}

// startScan will initiate the scan/parse process
func (req *MyRequest) startScan(l *log.Logger, base string) resultsapp.FinalResponse {

	if req.File == nil {
		resp, err := http.Get(req.URL)
		if err != nil {
			l.Println("Error fetching url response", err)
		}
		defer resp.Body.Close()

		//Call Parse Service to read http response body
		return req.parserEngine.Parse(resp.Body)

	} else {

		//Call Parse Service to read html file body
		return req.parserEngine.Parse(req.File)
	}
}
