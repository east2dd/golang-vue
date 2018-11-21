package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string
	Products []*Product `gorm:"many2many:categories_products;"`
}
