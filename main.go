package main

import (
	"webserver/resultsapp"
	"webserver/server"
)

var port string

func init() {
	//get env
	port = ":" + resultsapp.GetEnvValueFor("PORT")
}

func main() {
	server := server.NewAPIServer(port)
	server.Run()
}
