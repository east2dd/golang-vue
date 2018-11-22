package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	u "github.com/xyingsoft/golang-vue/utils"
)

type Category struct {
	gorm.Model
	Name     string     `json:"name"`
	Products []*Product `gorm:"many2many:categories_products;"`
}

func (category *Category) Validate() (map[string]interface{}, bool) {
	if category.Name == "" {
		return u.Message(false, "should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (category *Category) Create() map[string]interface{} {
	if resp, ok := category.Validate(); !ok {
		return resp
	}

	GetDB().Create(category)

	resp := u.Message(true, "success")
	resp["contact"] = category
	return resp
}

func GetCategory(id uint) *Category {
	category := &Category{}

	err := GetDB().Table("categories").Where("id = ?", id).First(category).Error
	if err != nil {
		return nil
	}
	return category
}

func GetCategories() []*Category {
	categories := make([]*Category, 0)

	err := GetDB().Table("categories").Find(&categories).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return categories
}

func GetProductsFor(category uint) []*Product {
	products := make([]*Product, 0)

	err := GetDB().Table("products").
		Joins("LEFT JOIN categories_products ON (products.id = categories_products.product_id)").
		Where("categories_products.category_id = ?", category).
		Find(&products).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}

func GetCategoriesFor(product uint) []*Category {
	categories := make([]*Category, 0)

	err := GetDB().Table("categories").
		Joins("LEFT JOIN categories_products ON (categories.id = categories_products.category_id)").
		Where("categories_products.product_id = ?", product).
		Find(&categories).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return categories
}
