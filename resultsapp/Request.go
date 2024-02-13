package resultsapp

import "io"

type MyRequest struct {
	URL   string    `json:"url" validate:"required,url"`
	Depth int       `json:"depth" validate:"gt=-1,lte=2"`
	File  io.Reader `json:"file"`
}

func NewMyRequest(URL string, depth int, file io.Reader) *MyRequest {
	return &MyRequest{URL: URL, Depth: depth, File: file}
}
