package seeders

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	// Seeder untuk Product Categories
	electronics := models.ProductCategory{Name: "Electronics"}
	fashion := models.ProductCategory{Name: "Fashion"}
	home := models.ProductCategory{Name: "Home & Living"}

	db.Create(&electronics)
	db.Create(&fashion)
	db.Create(&home)

	// Seeder untuk Users
	admin := models.User{Username: "admin", Email: "admin@gocommerce.com", Password: "admin123"}
	john := models.User{Username: "john_doe", Email: "john@example.com", Password: "johnpass"}
	jane := models.User{Username: "jane_doe", Email: "jane@example.com", Password: "janepass"}

	db.Create(&admin)
	db.Create(&john)
	db.Create(&jane)

	// Seeder untuk Products
	tv := models.Product{Name: "Smart TV", CategoryID: electronics.ID, Price: 3500000, Stock: 10}
	shirt := models.Product{Name: "T-Shirt", CategoryID: fashion.ID, Price: 75000, Stock: 50}
	sofa := models.Product{Name: "Sofa", CategoryID: home.ID, Price: 1500000, Stock: 5}

	db.Create(&tv)
	db.Create(&shirt)
	db.Create(&sofa)

	// Seeder untuk Transactions
	trans1 := models.Transaction{UserID: john.ID, Amount: tv.Price * 1}
	trans2 := models.Transaction{UserID: jane.ID, Amount: shirt.Price*2 + sofa.Price*1}

	db.Create(&trans1)
	db.Create(&trans2)

	// Seeder untuk Transaction Items
	item1 := models.TransactionItem{TransactionID: trans1.ID, ProductID: tv.ID, Quantity: 1}
	item2 := models.TransactionItem{TransactionID: trans2.ID, ProductID: shirt.ID, Quantity: 2}
	item3 := models.TransactionItem{TransactionID: trans2.ID, ProductID: sofa.ID, Quantity: 1}

	db.Create(&item1)
	db.Create(&item2)
	db.Create(&item3)
}
