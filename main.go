package main

import (
	"github.com/ahujatejas06/pclubtask5/controllers"
	"github.com/ahujatejas06/pclubtask5/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/", controllers.FindUsers)
	router.GET("/:id", controllers.FindUser)
	router.POST("/", controllers.CreateUser)
	router.PATCH("/:id", controllers.UpdateUser)
	router.DELETE("/:id", controllers.DeleteUser)

	// Run the server
	router.Run()
}
