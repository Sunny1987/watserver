package main

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webserver/server"
)

func main() {

	//log section
	l := log.New(os.Stdout, "WAT:", log.LstdFlags)

	//SetHandler
	routerHandler := server.GetNewLogger(l)

	//set mux
	serverMux := mux.NewRouter()

	//set env
	SetEnvVar()

	//Register the handlers to the server mux
	postRouter := serverMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/scan", routerHandler.GetURLResp)
	postRouter.HandleFunc("/uploadhtml", routerHandler.FileScan)
	postRouter.Use(routerHandler.MiddlewareValidation)

	//CORS
	ch := goHandlers.CORS(
		goHandlers.AllowedOrigins([]string{"*"}),
		goHandlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Origin"}),
		goHandlers.AllowedMethods([]string{"GET", "POST"}),
	)

	//Load the config.json properties
	//config, err := server.LoadConfiguration("config.json")
	//if err != nil {
	//
	//}
	port := os.Getenv("PORT")
	if port == "" {
		l.Fatal("Port must be set")
	}

	//Define server
	prodServer := &http.Server{
		Addr:         port,
		Handler:      ch(serverMux),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 50 * time.Minute,
		IdleTimeout:  50 * time.Minute,
		ErrorLog:     l,
	}

	go func() {
		myFigure := figure.NewFigure("WAT", "", true)
		myFigure.Print()
		l.Println("version: 1.0.0")
		l.Println("Author: Sabyasachi Roy")
		l.Println("Starting server...")
		if err := prodServer.ListenAndServe(); err != nil {
			l.Printf("Error starting server: %v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("Stopping server as per user interrupt", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := prodServer.Shutdown(tc)
	if err != nil {
		return
	}
}
