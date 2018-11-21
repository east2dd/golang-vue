package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	Price      int64
	Categories []*Category `gorm:"many2many:categories_products;"`
}

func (a *Product) Validate() bool {
	return true
}
