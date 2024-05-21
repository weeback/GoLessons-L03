package main

import (
	"fmt"
	"myapp"
)

var (
	bindAddr = "0.0.0.0:8080"
)

func main() {
	fmt.Printf("Server API\n------------------------------------\n"+
		"\tVersion: %s\n------------------------------------\n", myapp.Version)
}
