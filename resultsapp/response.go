package resultsapp

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

// ResponseBundle object for Response data management
type ResponseBundle struct {
	rw       http.ResponseWriter
	logger   *log.Logger
	response []FinalResponse
	printer  PrintService
}

// NewResponseBundle constructor for ResponseBundle object
func NewResponseBundle(rw http.ResponseWriter, logger *log.Logger, response []FinalResponse, printerService PrintService) *ResponseBundle {
	return &ResponseBundle{rw: rw, logger: logger, response: response, printer: printerService}
}

// FinalResponse object for response creation
type FinalResponse struct {
	Request *MyRequest  `json:"request"`
	Person  *string     `json:"person"`
	Results []TagResult `json:"results"`
	Doc     *html.Node  `json:"-"`
}

// PrintResponse will generate the response of the query
func (rBundle ResponseBundle) PrintResponse() {
	rBundle.logger.Println("Initiating the response....")

	var err error
	if len(rBundle.response) == 1 {
		if status, _ := rBundle.NoResultsGenerated(rBundle.response[0], err); status {
			return
		}
		if rBundle.ResultGenerated(rBundle.response[0]) {
			return
		}

	} else {
		var results []FinalResponse
		for i := 0; i < len(rBundle.response); i++ {
			respI := rBundle.response[i]
			if len(respI.Results) > 0 {
				results = append(results, respI)
				rBundle.printer.CreateHTMLPage(respI)
			}
		}
		if len(results) > 0 {
			rBundle.printer.CreateJSONAndPrintResponse(results)
		} else {
			rBundle.printer.CreateJSONAndPrintResponse("No WCAG errors observed for " + rBundle.response[0].Request.URL)
		}

	}
}

func (rBundle ResponseBundle) ResultGenerated(resp FinalResponse) bool {
	if len(resp.Results) == 1 {
		rBundle.printer.CreateJSONAndPrintResponse(resp)
		rBundle.printer.CreateHTMLPage(resp)
		return true
	}
	return false
}

func (rBundle ResponseBundle) NoResultsGenerated(resp FinalResponse, err error) (bool, string) {
	resp = rBundle.response[0]
	var entity string
	if len(resp.Results) == 0 {
		if resp.Request.URL != "" {
			_, err = fmt.Fprintf(rBundle.rw, "\n No WCAG compliance error observed for URL= %s", rBundle.response[0].Request.URL)
			entity = resp.Request.URL
		} else {
			_, err = fmt.Fprintf(rBundle.rw, "\n No WCAG compliance error observed for File= %s\n", rBundle.response[0].Request.FileName)
			entity = resp.Request.FileName
		}

		if err != nil {
			rBundle.logger.Printf("Error : %v", err)

		}
		return true, entity
	}
	return false, entity
}
