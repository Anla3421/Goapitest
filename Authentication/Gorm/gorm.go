package Gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Product struct {
// 	gorm.Model
// 	Code    string
// 	Price   uint
// 	Deleted gorm.DeletedAt
// }

type Product struct {
	gorm.Model
	ID    uint
	Code  string
	Price uint
	// Deleted gorm.DeletedAt //不使用gorm的方式啟用軟刪除
}

func Users() Product {
	db, err := gorm.Open(mysql.Open("root:adminstrator@tcp(127.0.0.1:3306)/testdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("successful connected")
	// // panic("failed to connect database")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	fmt.Println("successful Created Data")

	// Read
	var product Product
	// var aa Product

	// db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println("successful Updated Data")

	// Delete - delete product
	db.Where("price", "200").Delete(&product)
	// db.Take(&product, 3)
	db.Find(&product, []int{1, 2, 3})
	// db.Find(&product, 3)
	db.Find(&product, 3)
	fmt.Println("successful Delete Data")

	// return product
	return product
}
