package server

import (
	"log"
	"webserver/Database"
)

// The NewLogger is the Object for logging and RouterHandler wrapper
type NewLogger struct {
	myLogger *log.Logger
	db       Database.DBService
}

// GetNewLogger is the constructor for the NewLogger Object/Struct
func GetNewLogger(myLogger *log.Logger) RouterHandler {
	return &NewLogger{myLogger: myLogger}
}

// GetNewLoggerWithDB is the constructor for the NewLogger Object/Struct and DB object
func GetNewLoggerWithDB(logger *log.Logger, db Database.DBService) RouterHandler {
	return &NewLogger{myLogger: logger, db: db}
}
