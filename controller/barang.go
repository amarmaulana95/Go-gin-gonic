package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goapi.com/models"
)

type BarangInput struct {
	Id   string `json: "id_barang"`
	Nama string `json: "nama_barang"`
	Kode string `json: "kode"`
}

//GET DATA
func BarangTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//buat data array
	var data []models.Barang
	db.Find(&data)

	if len(data) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	} else {

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
