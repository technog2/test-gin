package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Status bool   `json:"status"`
}

func Index(ctx *gin.Context) {
	users := []User{
		User {
			Name: "Ali",
			Age: 20,
			Status: true,
		},
		User {
			Name: "Pedram",
			Age: 19,
			Status: false,
		},
	}

	ctx.JSON(http.StatusOK, gin.H {
		"users": users,
	})
}
