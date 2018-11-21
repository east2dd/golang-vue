package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xyingsoft/golang-vue/models"
)

func main() {
	db := models.GetDB()

	db.Debug().AutoMigrate(&models.Category{})
	db.Debug().AutoMigrate(&models.Product{})
	db.Debug().AutoMigrate(&models.Account{}) //Database migration

	db.Create(&models.Category{
		Name: "Mens",
		Products: []*models.Product{
			{Name: "Product 1"},
			{Name: "Product 2"},
		},
	})

	db.Create(&models.Category{
		Name: "Womens",
		Products: []*models.Product{
			{Name: "Product 1"},
			{Name: "Product 2"},
		},
	})
}
