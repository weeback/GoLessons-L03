package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myapp"
	"net/http"

	"myapp/api/gorilla-server"
)

var (
	bindAddr = "0.0.0.0:8080"
)

func main() {
	fmt.Printf("Server API\n------------------------------------\n"+
		"\tVersion: %s\n------------------------------------\n", myapp.Version)

	// Create a new router
	r := mux.NewRouter()

	// Register health-check handler
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "OK"); err != nil {
			log.Printf("%s - %s - Error writing response: %v", r.RemoteAddr, r.RequestURI, err)
			return
		}
	}).Methods(http.MethodGet)

	// Register your handlers
	gorilla_server.New().RegisterHandlers(r)

	// Walk through all the registered routes
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			fmt.Printf("[%-8s] %s\n", "", pathTemplate)
			return nil
		}
		for _, method := range methods {
			fmt.Printf("[%-8s] %s\n", method, pathTemplate)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking routes: ", err)
	}

	// Print the bind address
	fmt.Printf("Server listening on: %s ...\n", bindAddr)

	// Start the server
	if err := http.ListenAndServe(bindAddr, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
