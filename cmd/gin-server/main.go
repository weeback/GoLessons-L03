package main

import (
	"fmt"
	"myapp"
	gin_server "myapp/api/gin-server"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	bindAddr = "0.0.0.0:8080"
)

func main() {
	fmt.Printf("Server API\n------------------------------------\n"+
		"\tVersion: %s\n------------------------------------\n", myapp.Version)

	// Create a new router
	r := gin.New()

	// Register health-check handler
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Register your handlers
	gin_server.New().RegisterRouter(r)

	// Create the server configuration
	serv := http.Server{
		Addr:    bindAddr,
		Handler: r,
	}
	// Print the bind address
	fmt.Printf("Server listening on: %s ...\n", bindAddr)

	// Start the server
	if err := serv.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
