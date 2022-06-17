package main

import (
	"gin-apicrud/controllers"
	"gin-apicrud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run()
}
