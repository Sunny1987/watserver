package main

import "webserver/server"

var port string

func init() {
	//get env
	port = ":" + GetEnvValueFor("PORT")
}

func main() {
	server := server.NewAPIServer(port)
	server.Run()
}
