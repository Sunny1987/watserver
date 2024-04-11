package server

import "webserver/resultsapp"

const (
	DBTRUE  = "true"
	DBFALSE = "false"
)

var (
	ENV      = resultsapp.GetEnvValueFor("ENV")
	DEV_URL  = resultsapp.GetEnvValueFor("DEV_URL")
	PROD_URL = resultsapp.GetEnvValueFor("PROD_URL")
	USEDB    = resultsapp.GetEnvValueFor("DBFLAG")
)
