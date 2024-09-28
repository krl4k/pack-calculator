package http

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func NewRouter(handler *PackCalculatorHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/calculate", handler.CalculatePacks).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	return r
}
