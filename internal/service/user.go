package service

import "github.com/mythiee/golang-web-template/internal/model"

func GetUser() model.User {
	return model.User{
		Name: "zhangsan",
		Age:  18,
	}
}
