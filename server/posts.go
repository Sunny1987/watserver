package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
	"sync"
	"time"
	"webserver/analyzerapp"
	"webserver/parserapp/parser"
	"webserver/parserapp/sitemapbuilder"
	"webserver/resultsapp"
)

var (
	wg     sync.WaitGroup
	Client = resty.New()
)

const (
	DBTRUE  = "true"
	DBFALSE = "false"
)

// GetURLResp will scan the URL with desired depth and provide the accessibility results
func (l *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	l.myLogger.Println("GetURLResp called...")
	//track execution time for scan
	timeStart := time.Now()

	req := &MyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		l.myLogger.Println("Middleware: %v", err)
	}

	//get the list of links from sitemap
	links, base := sitemapbuilder.SiteMap(req.URL, req.Depth, l.myLogger)
	l.myLogger.Println("***** site map completed*****")
	var finalResult []resultsapp.FinalResponse

	//scan based on depth
	wg.Add(len(links))
	for i, link := range links {

		go func(link string) {
			defer wg.Done()

			//create requestBundle
			requestBundle := resultsapp.NewMyRequest(req.URL, req.Depth, req.File, req.FileName)

			//Create AnalyzerService
			analyzerEngine := analyzerapp.NewAnalyzeBundleNoCollectedTags(requestBundle, l.myLogger)

			//Create new ParserService
			parseEngine := parser.NewParseBundle(requestBundle, l.myLogger, base, analyzerEngine)

			//Request object created with constructor
			reqMod := NewMyRequestURL(link, req.Depth, parseEngine)

			l.myLogger.Printf("Link# %v : %v ", i, reqMod.URL)
			//start scan for url
			results := reqMod.startScan(l.myLogger, base)
			//mu.Lock()
			finalResult = append(finalResult, results)
			//mu.Unlock()

		}(link)
	}
	wg.Wait()

	//update the db if dbflag is set to true
	if useDB := resultsapp.GetEnvValueFor("DBFLAG"); useDB == DBTRUE {
		l.db.UpdateResults(req.Id, finalResult)
	}

	//print the response
	printer := resultsapp.NewPrinter(rw, l.myLogger)
	rBubdle := resultsapp.NewResponseBundle(rw, l.myLogger, finalResult, printer)
	rBubdle.PrintResponse()

	l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
}

// FileScan will scan the uploaded File and provide the accessibility results
func (l *NewLogger) FileScan(rw http.ResponseWriter, r *http.Request) {
	l.myLogger.Println("FileScan called...")
	//track execution time for scan
	timeStart := time.Now()

	var finalResult []resultsapp.FinalResponse

	file, handler, err := r.FormFile("file")

	if err != nil {
		l.myLogger.Println("Error Retrieving the File")
		l.myLogger.Println(err)
		return
	}
	defer file.Close()
	l.myLogger.Printf("Uploaded File: %+v\n", handler.Filename)
	l.myLogger.Printf("File Size: %+v\n", handler.Size)
	l.myLogger.Printf("MIME Headers: %+v\n", handler.Header)

	//create requestBundle
	requestBundle := resultsapp.NewMyRequest("", 0, file, handler.Filename)

	//Create AnalyzerService
	analyzerEngine := analyzerapp.NewAnalyzeBundleNoCollectedTags(requestBundle, l.myLogger)

	//Create new ParserService
	parseEngine := parser.NewParseBundle(requestBundle, l.myLogger, "", analyzerEngine)

	//Request object created with constructor
	reqMod := NewMyRequestFile(file, handler.Filename, parseEngine)

	//start scan for File
	results := reqMod.startScan(l.myLogger, "")

	//Create Final response
	finalResult = append(finalResult, results)

	//print the response
	printer := resultsapp.NewPrinter(rw, l.myLogger)
	rBubdle := resultsapp.NewResponseBundle(rw, l.myLogger, finalResult, printer)
	rBubdle.PrintResponse()

	l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
}

// ScanRegister will initiate scan and return an uuid to the user, this uuid will be used to fetch the scan results later
func (l *NewLogger) ScanRegister(writer http.ResponseWriter, request *http.Request) {
	l.myLogger.Println("ScanRegister called...")
	if useDB := resultsapp.GetEnvValueFor("DBFLAG"); useDB == DBFALSE {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("Database is disabled"))
		return
	}
	env := resultsapp.GetEnvValueFor("ENV")
	url := GetUrl(env)

	myreq := &MyRequest{}
	err := json.NewDecoder(request.Body).Decode(myreq)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		l.myLogger.Println("Error in retriving request body")
		return
	}

	recId := l.db.CreateResult(myreq.URL)
	myreq.Id = recId

	go func() {
		_, err = Client.R().SetHeader("Accept", "application/json").SetBody(myreq).Post(url)
		if err != nil {
			http.Error(writer, "Error creating POST request", http.StatusInternalServerError)
			return
		}
	}()

	resp := "Scan initiated with UUID=" + recId.String()
	writer.Write([]byte(resp))
}

func GetUrl(env string) string {
	var url string
	if env == "dev" {
		url = "http://localhost:8080/api/v1/scan"
	} else {
		url = resultsapp.GetEnvValueFor("PROD_URL") + "scan"
	}
	return url
}
