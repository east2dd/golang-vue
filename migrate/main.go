package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xyingsoft/golang-vue/config"
	"github.com/xyingsoft/golang-vue/data"
	"github.com/xyingsoft/golang-vue/model"
)

func main() {
	config.Init()
	data.Init()

	db := data.DB()
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Product{})

	db.Create(&model.Category{
		Name: "Mens",
		Products: []*model.Product{
			{Name: "Product 1", Price: 1000},
			{Name: "Product 2", Price: 2000},
		},
	})

	db.Create(&model.Category{
		Name: "Womens",
		Products: []*model.Product{
			{Name: "Product 1", Price: 1000},
			{Name: "Product 2", Price: 2000},
		},
	})

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)
}
