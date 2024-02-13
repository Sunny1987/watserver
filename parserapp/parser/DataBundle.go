package parser

import (
	"io"
	"webserver/resultsapp"
)

// DataBundle is the wrapper for ParseBundle
type DataBundle interface {
	Parse(responseBody io.Reader) resultsapp.Response
}
