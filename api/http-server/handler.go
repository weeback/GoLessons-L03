package http_server

import (
	"fmt"
	"log"
	"myapp"
	"myapp/pkg/net"
	"net/http"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) Register() {
	// Register the handler for the /api/version route
	// Check and handle the method, if allowed to then call the handler, otherwise return an error
	// Require the method to be GET
	http.HandleFunc("/api/v1/version", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// if method is allowed, call the handler
		h.getAppVersion(w, r)
	})

	// Use the WithMethods helper to handle the method check
	http.Handle("/api/v2/version", net.WithMethods(h.getAppVersion))
}

func (h *Handler) getAppVersion(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Version: %s", myapp.Version); err != nil {
		log.Printf("%s - %s - Error writing response: %v", r.RemoteAddr, r.RequestURI, err)
		return
	}
}
