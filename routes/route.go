package routes

import (
	"net/http"

	ctl "github.com/SitanayaIvan/latihan_golang/controller"
	"github.com/gin-gonic/gin"
)

func GetEndpoint(r *gin.Engine, ctl ctl.Controller) {
	v1 := r.Group("/api")
	//
	v1.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	v1.POST("/user", ctl.User.CreateUser)
	v1.GET("/user", ctl.User.GetUsers)
	v1.GET("/user/:id", ctl.User.GetUserById)
	v1.PUT("/user/:id", ctl.User.UpdateUser)
	v1.DELETE("/user/:id", ctl.User.DeleteUser)
}
