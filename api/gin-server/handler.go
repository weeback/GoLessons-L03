package gin_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myapp"
	"net/http"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct{}

func (h *Handler) RegisterRouter(r *gin.Engine) {
	r.GET("/api/version", h.getAppVersion)
}

func (h *Handler) getAppVersion(c *gin.Context) {
	// wrap the handler function to avoid the need to pass gin.Context
	func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Version: %s", myapp.Version); err != nil {
			log.Printf("%s - %s - Error writing response: %v", r.RemoteAddr, r.RequestURI, err)
			return
		}
	}(c.Writer, c.Request)
}
