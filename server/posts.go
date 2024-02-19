package server

import (
	"net/http"
	"sync"
	"time"
	"webserver/parserapp/sitemapbuilder"
	"webserver/resultsapp"
)

var wg sync.WaitGroup

func (l *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	req := r.Context().Value(KeyUser{}).(*MyRequest)

	//get the list of links from sitemap
	links, base := sitemapbuilder.SiteMap(req.URL, req.Depth, l.myLogger)
	l.myLogger.Println("***** site map completed*****")
	var finalResult []resultsapp.Response

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
	//PrintResponse(finalResult, rw, l.myLogger)

	l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
}

func (l *NewLogger) FileScan(rw http.ResponseWriter, r *http.Request) {

}
