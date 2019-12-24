package mappings

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/users/", controllers.GetUser)
		v1.POST("/login/", controllers.Login)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.POST("/users", controllers.PostUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}
	v2 := Router.Group("/v2")
	{
		v2.GET("/posts/", controllers.GetPosts)
		v2.GET("/posts/:postid", controllers.GetPostDetail)
		v2.POST("/posts/:userid", controllers.GetPostsDetailByUserID)

	}

}
