package server

import "net/http"

// PingServer will ping the server to check for server availability
func (l *NewLogger) PingServer(writer http.ResponseWriter, request *http.Request) {
	l.myLogger.Println("Ping called...")
	writer.Write([]byte("WATServer is live!"))
}

func (l *NewLogger) GetLatestResults(writer http.ResponseWriter, request *http.Request) {

}
