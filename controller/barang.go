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
	//find bawaan gorm  -> select *
	db.Find(&data)
	//cek data
	if len(data) <= 0 {
		//response
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func BarangAdd(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi
	var dataInput BarangInput
	// if yg di post kan struktur json bukan ?
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//buat variable untuk menamppung hasil inputan {data}
	data := models.Barang{
		Nama: dataInput.Nama,
		Kode: dataInput.Kode,
	}
	db.Create(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})

}

func BarangUpdate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi
	var data models.Barang
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak ditemukan"})
		return
	}
	var dataInput BarangInput
	// if yg di post kan struktur json bukan ?
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&data).Update(&dataInput)
	c.JSON(http.StatusOK, gin.H{"data": data})

}

func BarangDelete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi
	var data models.Barang
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak ditemukan"})
		return
	}
	db.Delete(&data)
	c.JSON(http.StatusOK, gin.H{"data": true})

}
