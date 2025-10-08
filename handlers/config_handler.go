package handlers

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/spotless_db?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&models.TransactionItem{}) // Migrate model yang diperlukan
}
