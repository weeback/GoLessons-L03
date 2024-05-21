package gorilla_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myapp"
	"net/http"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/version", h.getAppVersion).Methods("GET", "POST")
}

func (h *Handler) getAppVersion(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Version: %s", myapp.Version); err != nil {
		log.Printf("%s - %s - Error writing response: %v", r.RemoteAddr, r.RequestURI, err)
		return
	}
}
