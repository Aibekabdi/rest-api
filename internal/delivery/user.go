package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Aibekabdi/rest-api/internal/models"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/user" {
		w.WriteHeader(404)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	var text = models.UserModel{}
	if err := json.Unmarshal(body, &text); err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	id, err := h.services.User.PostUserindb(text)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	strid := strconv.Itoa(id)
	w.Write([]byte(strid))
}

func (h *Handler) userRUD(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[0:11] != "/rest/user/" || len(r.URL.Path) < 12 {
		w.WriteHeader(404)
		return
	}
	strId := r.URL.Path[11:]
	Id, _ := strconv.Atoi(strId)
	switch r.Method {
	case "GET":
		model, err := h.services.GetUser(Id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		res, err := json.Marshal(model)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	case "PUT":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		var text = models.UserModel{}
		if err := json.Unmarshal(body, &text); err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		if err := h.services.PutUser(Id, text); err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("Changed"))
	case "DELETE":
		err := h.services.DeleteUser(Id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Write([]byte("Deleted"))
		return
	default:
		w.WriteHeader(405)
		return
	}

}
