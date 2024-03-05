package resultsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	Request *MyRequest  `json:"request"`
	Person  *string     `json:"person"`
	Results []TagResult `json:"results"`
	Doc     *html.Node  `json:"-"`
}

// PrintResponse will generate the response of the query
func (rBundle ResponseBundle) PrintResponse() {
	rBundle.logger.Println("Initiating the response....")

	//Delete pre-existing text files
	DeleteTextFiles()

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
				CreateHTMLPage(rBundle.logger, respI)
			}
		}
		if len(results) > 0 {
			rBundle.CreateJSONAndPrintResponse(results)
		} else {

		}

	}
}

func (rBundle ResponseBundle) ResultGenerated(resp FinalResponse) bool {
	if len(resp.Results) == 1 {
		rBundle.CreateJSONAndPrintResponse(resp)
		CreateHTMLPage(rBundle.logger, resp)
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

// CreateJSONAndPrintResponse is responsible to convert the FinalResponse to JSON and print final response
func (rBundle ResponseBundle) CreateJSONAndPrintResponse(results interface{}) {
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

func CreatePDF(l *log.Logger, resp string, fileName string) {
	l.Println("...Generating PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, resp)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		l.Println(err)
	}
}

// CreateHTMLPage will generate a HTML text from the updated nodes
func CreateHTMLPage(l *log.Logger, resp FinalResponse) {
	l.Println(".. Generating HTML analysis file...")
	var buf bytes.Buffer
	doc := resp.Doc
	var fileN string

	if err := html.Render(&buf, doc); err != nil {
		log.Fatal(err)
	}
	if resp.Request.URL != "" {
		fileN = strings.Split(resp.Request.URL, "://")[1]
		fmt.Println(fileN)
	} else {
		fileN = strings.Split(resp.Request.FileName, ".")[0]
		fmt.Println(fileN)
	}

	fileN = fileN + "_analyzed.txt"

	// Print the reconstructed HTML body
	file, err := os.Create(fileN)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(file, buf.String())
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteTextFiles will delete all previous text files before the next scan completes
func DeleteTextFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".txt" {
			err := os.Remove(file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
