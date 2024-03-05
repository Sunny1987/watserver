package main

import "os"

// SetEnvVar function will set all necessary variables for the app
func SetEnvVar() {
	os.Setenv("PORT", ":8080")
}
