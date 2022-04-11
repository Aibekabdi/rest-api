package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/Aibekabdi/rest-api/internal/models"
)

func (h *Handler) counterAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	if len(r.URL.Path) < 19 {
		w.WriteHeader(404)
		return
	}
	if r.URL.Path[0:18] != "/rest/counter/add/" {
		w.WriteHeader(404)
		return
	}
	if err := h.services.Counter.Add(r.URL.Path[18:]); err != nil {
		w.Write([]byte(err.Error()))
	}
	return
}

func (h *Handler) counterSub(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	if len(r.URL.Path) < 19 {
		w.WriteHeader(404)
		return
	}
	if r.URL.Path[0:18] != "/rest/counter/sub/" {
		w.WriteHeader(404)
		return
	}
	if err := h.services.Counter.Sub(r.URL.Path[18:]); err != nil {
		w.Write([]byte(err.Error()))
	}
	return
}
func (h *Handler) counterVal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	if len(r.URL.Path) != 17 {
		w.WriteHeader(404)
		return
	}
	if r.URL.Path[0:17] != "/rest/counter/val" {
		w.WriteHeader(404)
		return
	}
	num, err := h.services.Counter.Val()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	counter := models.CounterModel{
		Num: num,
	}
	res, err := json.Marshal(counter)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
