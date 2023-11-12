package routes

import (
	"latihan_golang/controller"

	"github.com/gin-gonic/gin"
)

func GetEndpoints(r *gin.Engine, ctl *controller.Controller) {
	v1 := r.Group("/api")

	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	v1.POST("/user", ctl.User.CreateUser)
	v1.GET("/user", ctl.User.GetUsers)
	v1.GET("/user/:id", ctl.User.GetUserById)
	v1.PUT("/user/:id", ctl.User.UpdateUser)
	v1.DELETE("/user/:id", ctl.User.DeleteUser)
}
