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
	URL   string    `json:"url" validate:"required,url"`
	Depth int       `json:"depth" validate:"gt=-1,lte=2"`
	File  io.Reader `json:"file"`
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
func (req *MyRequest) startScan(l *log.Logger, base string) resultsapp.Response {

	//create requestBundle
	requestBundle := resultsapp.NewMyRequest(req.URL, req.Depth, req.File)

	//Create new ParseBundle/DataBundle
	DataBundle := parser.NewParseBundle(requestBundle, l, base)

	if req.File == nil {
		resp, err := http.Get(req.URL)
		if err != nil {
			l.Println("Error fetching url response", err)
		}
		defer resp.Body.Close()

		//Call Parse Method to read http response body
		return DataBundle.Parse(resp.Body)

	} else {

		//Call Parse Method to read html file body
		return DataBundle.Parse(req.File)
	}
}
