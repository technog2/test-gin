package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H {
		"message": "Hello " + name,
	})
}
