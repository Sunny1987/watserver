package main

import (
	"webserver/resultsapp"
	"webserver/server"
)

var port string
var dns string

// init function takes in the port and DB details
func init() {
	//get env
	port = ":" + resultsapp.GetEnvValueFor("PORT")
	dns = resultsapp.GetEnvValueFor("DNS_PGEDGE")
}

func main() {
	//declare server
	server := server.NewAPIServer(port, dns)

	//start the server execution
	if err := server.Run(); err != nil {
		panic(err)
	}
}
