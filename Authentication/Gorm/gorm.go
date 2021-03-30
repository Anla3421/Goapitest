package Gorm

import (
	"fmt"
	"time"

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

type Action_log struct {
	ID          uint `gorm:"primarykey"`
	User        uint
	Url         string
	Action      string `sql:"enum('Read', 'Update', 'Create')"`
	Result      string `sql:"enum('success', 'fail')"`
	Origin_data string
	Alter_data  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Deleted gorm.DeletedAt //不使用gorm的方式啟用軟刪除

}

func (Action_log) TableName() string {
	return "action_log"
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

func Actionlogapi() []*Action_log {
	db, err := gorm.Open(mysql.Open("root:adminstrator@tcp(127.0.0.1:3306)/jared?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("successful connected")
	var actionlogTable []*Action_log
	// result := db.Table("action_log").Find("action_log")
	// result.RowsAffected
	db.Table("action_log").Find(&actionlogTable)
	return actionlogTable
}
