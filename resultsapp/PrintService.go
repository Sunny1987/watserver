package resultsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// PrintService contracts
type PrintService interface {
	CreateHTMLPage(resp FinalResponse)
	CreateJSONAndPrintResponse(results interface{})
	CreatePDF(resp string, fileName string)
}

// Printer object
type Printer struct {
	rw     http.ResponseWriter
	logger *log.Logger
}

// NewPrinter is the constructor for Printer
func NewPrinter(rw http.ResponseWriter, logger *log.Logger) PrintService {
	return Printer{rw: rw, logger: logger}
}

// CreateJSONAndPrintResponse is responsible to convert the FinalResponse to JSON and print final response
func (printer Printer) CreateJSONAndPrintResponse(results interface{}) {
	rep, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		printer.logger.Println(err)
	}

	finalResponse := string(rep)
	_, err = fmt.Fprintln(printer.rw, finalResponse)
	if err != nil {
		printer.logger.Printf("Error : %v", err)
	}

}

// CreateHTMLPage will generate a HTML text from the updated nodes
func (printer Printer) CreateHTMLPage(resp FinalResponse) {
	printer.logger.Println(".. Generating HTML analysis file...")
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

	// Set headers for file download
	printer.rw.Header().Set("Content-Disposition", "attachment; filename=html_to_txt.txt")
	printer.rw.Header().Set("Content-Type", "text/plain")

	// Write the HTML content to response
	if _, err := printer.rw.Write(buf.Bytes()); err != nil {
		log.Println("Failed to write response:", err)
	}

	// Print the reconstructed HTML body

	////created directory
	//dirPath := createDirectory(printer.logger)
	//
	////Delete old files
	//deleteTextFiles(dirPath)

	//define filepath+

	//filePath := dirPath + "/" + fileN
	//file, err := os.Create(filePath)
	//defer file.Close()
	//if err != nil {
	//	printer.logger.Fatalf("Failed to create file: %v", err)
	//}

	//_, err = io.WriteString(file, buf.String())
	//if err != nil {
	//	printer.logger.Fatal(err)
	//}
	//printer.logger.Printf("Created file: %s", filePath)
}

func createDirectory(l *log.Logger) string {
	dirPath := "C:/Results"
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		l.Fatalf("Failed to create directory: %v", err)
	}
	l.Printf("Created directory: %s", dirPath)
	return dirPath
}

// deleteTextFiles will delete all previous text files before the next scan completes
func deleteTextFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if filepath.Ext(name) == ".txt" {
			err := os.Remove(filepath.Join(path, name))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (printer Printer) CreatePDF(resp string, fileName string) {
	printer.logger.Println("...Generating PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, resp)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		printer.logger.Println(err)
	}
}
