package server

import (
	"net/http"
	"sync"
	"time"
	"webserver/parserapp/sitemapbuilder"
	"webserver/resultsapp"
)

var wg sync.WaitGroup

// GetURLResp will scan the URL with desired depth and provide the accessibility results
func (l *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	req := r.Context().Value(KeyUser{}).(*MyRequest)

	//get the list of links from sitemap
	links, base := sitemapbuilder.SiteMap(req.URL, req.Depth, l.myLogger)
	l.myLogger.Println("***** site map completed*****")
	var finalResult []resultsapp.FinalResponse

	//scan based on depth
	wg.Add(len(links))
	for i, link := range links {

		go func(link string) {
			defer wg.Done()

			//setup req object
			reqMod := &MyRequest{}
			reqMod.URL = link
			reqMod.Depth = req.Depth

			l.myLogger.Printf("Link# %v : %v ", i, reqMod.URL)
			//start scan for url
			results := reqMod.startScan(l.myLogger, base)
			//mu.Lock()
			finalResult = append(finalResult, results)
			//mu.Unlock()

		}(link)

	}

	wg.Wait()

	//print the response
	resultsapp.PrintResponse(rw, l.myLogger, finalResult)

	l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
}

// FileScan will scan the uploaded File and provide the accessibility results
func (l *NewLogger) FileScan(rw http.ResponseWriter,
	r *http.Request) {
	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	//req := r.Context().Value(KeyUser{}).(*MyRequest)

	var finalResult []resultsapp.FinalResponse

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(rw, "Error max file size exceeded", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("myfile")

	if err != nil {
		l.myLogger.Println("Error Retrieving the File")
		l.myLogger.Println(err)
		return
	}
	defer file.Close()
	l.myLogger.Printf("Uploaded File: %+v\n", handler.Filename)
	l.myLogger.Printf("File Size: %+v\n", handler.Size)
	l.myLogger.Printf("MIME Headers: %+v\n", handler.Header)

	//setup req object
	reqMod := &MyRequest{}
	reqMod.URL = ""
	reqMod.Depth = 0
	reqMod.File = file
	reqMod.FileName = handler.Filename

	//start scan for File
	results := reqMod.startScan(l.myLogger, "")

	//Create Final response
	finalResult = append(finalResult, results)

	//print the response
	resultsapp.PrintResponse(rw, l.myLogger, finalResult)

	l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
}
