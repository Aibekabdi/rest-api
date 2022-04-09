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
	//rest
	mux.HandleFunc("/rest", h.rest)
	//substr
	mux.HandleFunc("/rest/substr", h.substr)
	mux.HandleFunc("/rest/substr/find", h.substrFind)

	return mux
}
