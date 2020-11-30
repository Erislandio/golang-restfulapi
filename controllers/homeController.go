package controllers

import (
	"github.com/eris/redis-example/utils"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	utils.ToJson(w, struct {
		Message string `json:"message"`
	}{
		Message: "Hello from utils",
	})
}
