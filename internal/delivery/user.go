package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

	var text = models.User{}
	if err := json.Unmarshal(body, &text); err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handler) userRUD(w http.ResponseWriter, r *http.Request) {

}
