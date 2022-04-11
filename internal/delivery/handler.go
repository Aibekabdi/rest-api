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
	//substr
	mux.HandleFunc("/rest/substr/find", h.substrFind)
	//email
	mux.HandleFunc("/rest/email/check", h.emailCheck)
	//counter
	mux.HandleFunc("/rest/counter/add/", h.counterAdd)
	mux.HandleFunc("/rest/counter/sub/", h.counterSub)
	mux.HandleFunc("/rest/counter/val", h.counterVal)
	//user
	mux.HandleFunc("/rest/user", h.createUser)
	mux.HandleFunc("/rest/user/", h.userRUD) //read-update-delete
	return mux
}
