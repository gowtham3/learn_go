package main

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"

	"example.com/m/v2/models" // new
)

func main() {
	r := gin.Default()

	models.ConnectDataBase() // new

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
