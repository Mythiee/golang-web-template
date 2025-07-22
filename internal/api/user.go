package api

import (
	"encoding/json"
	"net/http"

	HttpConstant "github.com/mythiee/golang-web-template/internal/constant/http"
	"github.com/mythiee/golang-web-template/internal/service"
	"github.com/mythiee/golang-web-template/types"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := service.GetUser()

	response := types.Response{
		Code:    HttpConstant.OK,
		Message: "success",
		Data:    user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HttpConstant.OK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), HttpConstant.INTERNAL_SERVER_ERROR)
	}
}
