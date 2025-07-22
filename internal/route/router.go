package route

import (
	"net/http"

	"github.com/mythiee/golang-web-template/internal/api"
)

func RegisterRoutes() {
	http.HandleFunc("/", api.GetUser)
}
