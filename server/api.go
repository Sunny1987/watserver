package server

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webserver/Database"
	"webserver/resultsapp"
)

type APIServer struct {
	addr string
	dns  string
}

func NewAPIServer(addr string, dns string) *APIServer {
	return &APIServer{addr: addr, dns: dns}
}

func (api *APIServer) Run() error {
	var err error

	//log section
	l := log.New(os.Stdout, "WAT:", log.LstdFlags)

	//initiate database
	var routerHandler RouterHandler
	if useDB := resultsapp.GetEnvValueFor("DBFLAG"); useDB == DBTRUE {
		l.Printf("DB Option=%v", DBTRUE)
		dBundle := Database.NewDBBundle(l)

		if err := dBundle.InitDB(api.dns); err != nil {
			return err
		}
		l.Println("DB connected successfully")

		routerHandler = GetNewLoggerWithDB(l, dBundle)
	} else {
		l.Printf("DB Option=%v", DBFALSE)
		routerHandler = GetNewLogger(l)
	}

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("POST /scan", routerHandler.GetURLResp)
	serverMux.HandleFunc("POST /uploadhtml", routerHandler.FileScan)
	serverMux.HandleFunc("POST /scanregister", routerHandler.ScanRegister)
	serverMux.HandleFunc("GET /ping", routerHandler.PingServer)
	serverMux.HandleFunc("GET /results/{uuid}", routerHandler.GetLatestResults)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", serverMux))

	//middleware connect
	middlewareChain := MiddlewareChain(
		routerHandler.MiddlewareValidationForScanRegister,
		routerHandler.MiddlewareValidationForScan,
		routerHandler.MiddlewareValidationForFile,
		routerHandler.MiddlewareForCorsUpdate,
	)

	//Define server
	prodServer := &http.Server{
		Addr:         api.addr,
		Handler:      middlewareChain(v1),
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
		l.Printf("Starting server at port %v", api.addr)
		if err = prodServer.ListenAndServe(); err != nil {
			l.Printf("Error starting server %v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("Stopping server as per user interrupt", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = prodServer.Shutdown(tc)
	if err != nil {
		l.Println(err)
		return err
	}
	return err
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}

}
