package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi.com/controller"
	"goapi.com/models"
)

func main() {
	r := gin.Default()

	//panggil models
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Rest Api test"})
	})

	r.GET("/barang", controller.BarangTampil)

	r.Run()
}
