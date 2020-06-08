package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/blrobin2/go-local-api/models"
	"github.com/blrobin2/go-local-api/controllers"
)

func main() {
	r := gin.Default()

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