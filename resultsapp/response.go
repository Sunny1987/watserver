package resultsapp

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"net/http"
)

// ResponseBundle object for Response data management
type ResponseBundle struct {
	rw       http.ResponseWriter
	logger   *log.Logger
	response []FinalResponse
}

// NewResponseBundle constructor for ResponseBundle object
func NewResponseBundle(rw http.ResponseWriter, logger *log.Logger, response []FinalResponse) *ResponseBundle {
	return &ResponseBundle{rw: rw, logger: logger, response: response}
}

// FinalResponse object for response creation
type FinalResponse struct {
	Request interface{} `json:"request"`
	Person  *string     `json:"person"`
	Results []TagResult `json:"results"`
}

// PrintResponse will generate the response of the query
func (rBundle ResponseBundle) PrintResponse() {
	rBundle.logger.Println("Initiating the response....")

	if len(rBundle.response) == 0 {
		_, err := fmt.Fprintln(rBundle.rw, "no bytes to unmarshal")
		if err != nil {
			rBundle.logger.Printf("Error : %v", err)

		}
		return
	}

	if len(rBundle.response) == 1 {
		var resp FinalResponse
		resp = rBundle.response[0]
		rBundle.CreateJSONandPrintResponse(resp)
		return
	}

	rBundle.CreateJSONandPrintResponse(rBundle.response)
}

func (rBundle ResponseBundle) CreateJSONandPrintResponse(results interface{}) {
	rep, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		rBundle.logger.Println(err)
	}

	finalResponse := string(rep)
	_, err = fmt.Fprintln(rBundle.rw, finalResponse)
	if err != nil {
		rBundle.logger.Printf("Error : %v", err)
	}
}

func CreatePDF(l *log.Logger, resp string) {
	l.Println("...Generating PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, resp)
	err := pdf.OutputFileAndClose("Report.pdf")
	if err != nil {
		l.Println(err)
	}
}
