package resultsapp

import "io"

type MyRequest struct {
	URL      string    `json:"url" validate:"required,url"`
	Depth    int       `json:"depth" validate:"gt=-1,lte=2"`
	File     io.Reader `json:"file"`
	FileName string    `json:"fileName"`
}

//func NewMyRequest(URL string, depth int, fileName string) *MyRequest {
//	return &MyRequest{URL: URL, Depth: depth, FileName: fileName}
//}

func NewMyRequest(URL string, depth int, file io.Reader, fileName string) *MyRequest {
	return &MyRequest{URL: URL, Depth: depth, File: file, FileName: fileName}
}
