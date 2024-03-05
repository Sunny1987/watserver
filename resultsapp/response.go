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
	var err error
	if len(rBundle.response[0].Results) == 0 {
		if rBundle.response[0].Request.URL != "" {
			_, err = fmt.Fprintf(rBundle.rw, "\n No WCAG compliance error observed for URL= %s", rBundle.response[0].Request.URL)

		} else {
			_, err = fmt.Fprintf(rBundle.rw, "\n No WCAG compliance error observed for File= %s\n", rBundle.response[0].Request.FileName)
		}
		if err != nil {
			rBundle.logger.Printf("Error : %v", err)

		}
		return
	}

	if len(rBundle.response[0].Results) == 1 {
		var resp FinalResponse
		resp = rBundle.response[0]
		rBundle.CreateJSONAndPrintResponse(resp)
		CreateHTMLPage(rBundle.logger, resp)
		return
	}

	for i := 0; i < len(rBundle.response); i++ {
		respI := rBundle.response[i]
		if len(respI.Results) > 0 {
			rBundle.CreateJSONAndPrintResponse(respI)
			CreateHTMLPage(rBundle.logger, respI)
		}
	}
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

	//Delete pre-existing text files
	DeleteTextFiles()

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
