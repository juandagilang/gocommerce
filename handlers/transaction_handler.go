package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, transaction)
	}
}

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mulai transaksi
		tx := db.Begin()
		if tx.Error != nil {
			c.JSON(500, gin.H{"message": "Failed to start transaction"})
			return
		}

		// Pastikan transaksi di-rollback jika terjadi panic
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		var input models.Transaction
		if err := c.ShouldBindJSON(&input); err != nil {
			tx.Rollback()
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		// Validasi setiap item untuk memastikan produk tersedia
		for _, item := range input.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				tx.Rollback()
				c.JSON(400, gin.H{"message": "Invalid product ID"})
				return
			}
		}

		// Buat transaksi dan item-itemnya
		if err := tx.Create(&input).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "Failed to create transaction"})
			return
		}

		// Commit jika semua operasi sukses
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "Failed to commit transaction"})
			return
		}

		c.JSON(201, input)
	}
}
