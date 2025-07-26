package main

import (
	"fmt"
	"net/http"

	"github.com/mythiee/golang-web-template/config/db"
	"github.com/mythiee/golang-web-template/config/env"
	"github.com/mythiee/golang-web-template/internal/route"
)

func main() {
	env.InitEnv()
	db.Init()
	route.RegisterRoutes()

	fmt.Println("server start on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
