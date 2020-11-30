package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Home).Methods(http.MethodGet)
	return r
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_ = json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	})
}