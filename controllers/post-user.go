package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"name":  user.Name,
		"email": user.Email,
	})
}
