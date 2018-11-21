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
}
