package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/blrobin2/go-local-api/api/controllers"
	"github.com/blrobin2/go-local-api/api/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/contracts", controllers.FindContracts)
	r.POST("/contracts", controllers.CreateContract)
	r.GET("/contracts/:id", controllers.FindContract)
	r.PATCH("/contracts/:id", controllers.UpdateContract)
	r.DELETE("/contracts/:id", controllers.DeleteContract)

	r.Run()
}