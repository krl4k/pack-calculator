package http

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func NewRouter(handler *PackCalculatorHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/calculate", handler.CalculatePacks).Methods("GET")
	r.HandleFunc("/api/pack-sizes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetPackSizes(w, r)
		case http.MethodPut:
			handler.UpdatePackSizes(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	return r
}
