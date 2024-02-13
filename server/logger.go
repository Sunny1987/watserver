package server

import "log"

// The NewLogger is the Object for logging and RouterHandler wrapper
type NewLogger struct {
	myLogger *log.Logger
}

// GetNewLogger is the constructor for the NewLogger Object/Struct
func GetNewLogger(logger *log.Logger) RouterHandler {
	return &NewLogger{myLogger: logger}
}
