package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	u "github.com/xyingsoft/golang-vue/utils"
)

type Product struct {
	gorm.Model
	Name       string      `json:"name"`
	Categories []*Category `gorm:"many2many:categories_products;"`
}

func (product *Product) Validate() (map[string]interface{}, bool) {
	if product.Name == "" {
		return u.Message(false, "should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (product *Product) Create() map[string]interface{} {
	if resp, ok := product.Validate(); !ok {
		return resp
	}

	GetDB().Create(product)

	resp := u.Message(true, "success")
	resp["contact"] = product
	return resp
}

func GetProduct(id uint) *Product {
	product := &Product{}

	err := GetDB().Table("categories").Where("id = ?", id).First(product).Error

	if err != nil {
		return nil
	}
	return product
}

func GetProducts() []*Product {
	products := make([]*Product, 0)

	err := GetDB().Table("products").Find(&products).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}
