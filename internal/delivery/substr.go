package delivery

import (
	"net/http"
	"text/template"
)

func (h *Handler) substr(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/substr" {
		w.WriteHeader(400)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	tmpl, err := template.ParseFiles("./template/substr.html")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(500)
		return
	}
}

func (h *Handler) substrFind(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/substr/find" {
		w.WriteHeader(400)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

}
