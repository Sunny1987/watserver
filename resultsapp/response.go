package resultsapp

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"net/http"
	"reflect"
)

type FinalResponse struct {
	Request interface{}
	Person  *string
	Results []TagResult
}

func PrintResponse(rw http.ResponseWriter, l *log.Logger, results []FinalResponse) {
	l.Println("Initiating the response....")
	var finalResponse string
	if len(results) == 0 {
		_, err := fmt.Fprintln(rw, "no bytes to unmarshal")
		if err != nil {
			l.Printf("Error : %v", err)
		}
	}

	rep, err := json.MarshalIndent(results, "", " ")
	//rep, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		l.Println(err)
	}

	finalResponse = string(rep)
	_, err = fmt.Fprintln(rw, finalResponse)
	if err != nil {
		l.Printf("Error : %v", err)
	}
}

//func ParseResultsAndCreatePDF(l *log.Logger, results []Response) string {
//	l.Println(".... Processing Results for PDF... ")
//	var finalResponse string
//	if len(results) == 0 {
//		_, err := fmt.Fprintln(rw, "no bytes to unmarshal")
//		if err != nil {
//			l.Printf("Error : %v", err)
//		}
//	}
//	return finalResponse
//}

func CreatePDF(l *log.Logger, resp string) {
	l.Println("...Generating PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, resp)
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		l.Println(err)
	}
}

func (resp FinalResponse) GetFilteredResponse() ([]string, map[string]interface{}) {
	tags := make(map[string]interface{})
	var tagNames []string

	structVal := reflect.ValueOf(resp)
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		name := structVal.Type().Field(i).Name
		value := field.Interface()

		if !reflect.ValueOf(value).IsNil() {
			tags[name] = value
			tagNames = append(tagNames, name)
		}
	}
	return tagNames, tags
}
