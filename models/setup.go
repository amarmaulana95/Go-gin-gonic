package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	//setup koneksi db
	db, err := gorm.Open("mysql", "root:@(localhost)/db_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal konek db")
	}
	db.AutoMigrate(&Barang{})
	return db
}
