package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Aibekabdi/rest-api/internal/models"
)

func (h *Handler) substrFind(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/substr/find" {
		w.WriteHeader(404)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var text = models.SubstrModel{}
	if err := json.Unmarshal(body, &text); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	text.Text = h.services.Substr.LongestSubstrFind(text.Text)
	res, err := json.Marshal(text)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
