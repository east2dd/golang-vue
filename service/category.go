package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/xyingsoft/golang-vue/data"
	"github.com/xyingsoft/golang-vue/model"
)

type CategoryParams struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func (m *CategoryParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&m.Name: "name",
	}
}

func GetCategories() (categories []model.Category, err error) {
	fmt.Println("Get products service")
	data.DB().Table("products").
		Select(`
			name,
			price
		`).Order("name ASC").Scan(&categories)

	if len(categories) == 0 {
		fmt.Println("Unable to find any categories!")
		err = errors.New("Unable to find any categories!")
	}

	return categories, err
}
