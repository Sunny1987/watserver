package server

import (
	"github.com/Sunny1987/ServerBase/server"
	"log"
	"os"
	"webserver/Database"
)

// APIServer is the wrapper over server functions
type APIServer struct {
	addr string
	dns  string
}

// NewAPIServer is the constructor for APIServer
func NewAPIServer(addr string, dns string) *APIServer {
	return &APIServer{addr: addr, dns: dns}
}

// Run initiates server execution
func (api *APIServer) Run() error {
	var err error
	//log section
	l := log.New(os.Stdout, "WAT:", log.LstdFlags)

	//initiate database
	var routerHandler RouterHandler

	if USEDB == DBTRUE {
		l.Printf("DB Option=%v", DBTRUE)
		dBundle := Database.NewDBBundle(l)

		if err = dBundle.InitDB(api.dns); err != nil {
			return err
		}
		l.Println("DB connected successfully")

		routerHandler = GetNewLoggerWithDB(l, dBundle)

	} else {
		l.Printf("DB Option=%v", DBFALSE)
		routerHandler = GetNewLogger(l)
	}

	// Create the app configuration
	app := server.NewMyAPIServer(&server.OptionalParams{
		Addr:      api.addr,
		AppName:   "WAT",
		AppAuthor: "Sabyasachi Roy",
		AppVer:    "1.0.0",
	})

	//Register Post calls
	app.Post("/scan", routerHandler.GetURLResp)
	app.Post("/uploadhtml", routerHandler.FileScan)
	app.Post("/scanregister", routerHandler.ScanRegister)

	//Register Get Calls
	app.Get("/ping", routerHandler.PingServer)
	app.Get("/results", routerHandler.GetLatestResults)
	app.Get("/results/{id}", routerHandler.GetResult)

	//SetUp Prefix
	app.AddPrefix("/api/v1/")

	//Add Middleware
	app.AddMiddleware(routerHandler.MiddlewareValidationForScanRegister)
	app.AddMiddleware(routerHandler.MiddlewareValidationForScan)
	app.AddMiddleware(routerHandler.MiddlewareValidationForFile)
	app.AddMiddleware(routerHandler.MiddlewareForResults)
	app.AddMiddleware(routerHandler.MiddlewareForResult)
	app.AddMiddleware(routerHandler.MiddlewareForCorsUpdate)

	//Run server
	if err = app.Run(); err != nil {
		return err
	}
	return err
}
