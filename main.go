package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/technog2/test-gin/router"

	// "github.com/technog2/test-gin/controllers/index"
	// "github.com/technog2/test-gin/controllers/ping"
	// "github.com/technog2/test-gin/controllers/get-user"
	// "github.com/technog2/test-gin/controllers/post-user"
	"github.com/technog2/test-gin/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	// apiv1.Use(jwt.JWT())

	{
		apiv1.GET("/", controller.Index)
		apiv1.GET("/ping", controller.Ping)
		apiv1.GET("/user/:name", controller.GetUser)
		apiv1.POST("/user", controller.PostUser)
	}

	return r
}

func main() {
	// router := gin.Default()
	r := InitRouter()

	r.Run(":4000") // listen and serve on localhost:4000
}














// package main

// import (
// 	"fmt"
// 	"net/http"
// 	// "io/ioutil"
// 	// "log"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	Name   string `json:"name"`
// 	Age    int8   `json:"age"`
// 	Status bool   `json:"status"`
// }

// func main() {
// 	router := gin.Default()

// 	// router.LoadHTMLGlob("templates/*")
// 	// //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

// 	// router.GET("/index", func(c *gin.Context) {
// 	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
// 	// 		"title": "Main website",
// 	// 	})
// 	// })

// 	router.GET("/index", func(c *gin.Context) {
// 		users := []User{
// 			User {
// 				Name: "Ali",
// 				Age: 20,
// 				Status: true,
// 			},
// 			User {
// 				Name: "Pedram",
// 				Age: 19,
// 				Status: false,
// 			},
// 		}

// 		c.JSON(http.StatusOK, gin.H {
// 			"users": users,
// 		})
// 	})

// 	router.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H {
// 			"message": "pong",
// 		})
// 	})

// 	router.GET("/user/:name", func(ctx *gin.Context) {
// 		name := ctx.Param("name")

// 		// ctx.String(200, "Hello %s", name)
// 		ctx.JSON(http.StatusOK, gin.H {
// 			"message": "Hello " + name,
// 		})
// 	})

// 	router.POST("/user", func(ctx *gin.Context) {
// 		var user struct {
// 			Name  string `json:"name"`
// 			Email string `json:"email"`
// 		}

// 		if err := ctx.ShouldBind(&user); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H {
// 			"name":  user.Name,
// 			"email": user.Email,
// 		})
// 	})


// 	// router.LoadHTMLGlob("templates/**/*")

// 	// // router.GET("/tailwindcss", func(c *gin.Context) { c.File("js/tailwindcss.js") })
// 	// router.GET("/tailwindcss", func(c *gin.Context) { 
// 	// 	content := ioutil.ReadFile("templates/js/tailwindcss.js")
// 	// 	c.Header("Content-Type", "text/javascript")
// 	// 	c.Data(200, "text/javascript", content)
// 	// })

// 	// router.GET("/posts/index", func(c *gin.Context) {
// 	// 	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
// 	// 		"title": "Posts",
// 	// 	})
// 	// })

// 	// router.GET("/users/index", func(c *gin.Context) {
// 	// 	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
// 	// 		"title": "Users",
// 	// 	})
// 	// })

// 	router.Run(":4000") // listen and serve on 0.0.0.0:8080

// 	fmt.Println("Running on port 4000")
// }
