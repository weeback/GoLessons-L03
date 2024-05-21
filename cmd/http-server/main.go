package main

import (
	"fmt"
	"myapp"
	"myapp/api/http-server"
	"net/http"
	"os"
)

var (
	bindAddr = "0.0.0.0:8080"
)

func main() {
	fmt.Printf("Server API\n------------------------------------\n"+
		"\tVersion: %s\n------------------------------------\n", myapp.Version)

	// Create a new handler
	http_server.New().Register()

	// Print the bind address
	fmt.Printf("Server listening on: %s ...\n", bindAddr)

	// Start the server
	if err := http.ListenAndServe(bindAddr, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
