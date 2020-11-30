package routes

import (
	"github.com/eris/redis-example/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods(http.MethodGet)
	return r
}
