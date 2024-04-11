package server

import (
	"net/http"
	"strings"
	"time"
)

// PingServer will ping the server to check for server availability
func (l *NewLogger) PingServer(writer http.ResponseWriter, request *http.Request) {
	l.myLogger.Println("Ping called...")
	writer.Write([]byte("WATServer is live!"))
}

func (l *NewLogger) GetLatestResults(writer http.ResponseWriter, request *http.Request) {
	l.myLogger.Println("GetLatestResults called...")
	//track execution time for scan
	timeStart := time.Now()
	if USEDB == DBTRUE {
		resp, err := l.db.GetRecordsForScan()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(resp))
		l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
	} else {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("DB access is not allowed"))
		return
	}
}

func (l *NewLogger) GetResult(writer http.ResponseWriter, request *http.Request) {
	l.myLogger.Println("Ping called...")

	//track execution time for scan
	timeStart := time.Now()
	if USEDB == DBTRUE {
		id := strings.Split(request.URL.Path, "/")[2]
		resp, err := l.db.GetRecordForId(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(resp))
		l.myLogger.Printf("Query completed in %v\n", time.Since(timeStart))
	} else {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("DB access is not allowed"))
		return
	}
}
