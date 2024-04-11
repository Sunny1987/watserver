package main

import (
	"webserver/resultsapp"
	"webserver/server"
)

var port string
var dns string

func init() {
	//get env
	port = ":" + resultsapp.GetEnvValueFor("PORT")
	dns = resultsapp.GetEnvValueFor("DNS_PGEDGE")
}

func main() {
	server := server.NewAPIServer(port, dns)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
