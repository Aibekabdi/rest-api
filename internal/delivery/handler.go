package delivery

import (
	"net/http"

	"github.com/Aibekabdi/rest-api/internal/service"
)

type Handler struct {
	services *service.Servise
}

func NewHandler(services *service.Servise) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	//view
	mux.HandleFunc("/", h.home)
	//substr
	mux.HandleFunc("/rest/substr/find", h.substrFind)

	return mux
}
