package router

import (
	"github.com/gin-gonic/gin"

	"github.com/technog2/test-gin/controllers/index"
	"github.com/technog2/test-gin/controllers/ping"
	"github.com/technog2/test-gin/controllers/get-user"
	"github.com/technog2/test-gin/controllers/post-user"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	// apiv1.Use(jwt.JWT())

	{
		apiv1.GET("/", Index)
		apiv1.GET("/ping", Ping)
		apiv1.GET("/user/:name", GetUser)
		apiv1.POST("/user", PostUser)
	}

	return r
}
