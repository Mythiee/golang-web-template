package api

import (
	"encoding/json"
	"net/http"

	"github.com/mythiee/golang-web-template/internal/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := service.GetUser()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
