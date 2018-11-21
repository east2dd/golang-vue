package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/xyingsoft/golang-vue/data"
	"github.com/xyingsoft/golang-vue/model"
)

type ProductParams struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func (m *ProductParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&m.Name:  "name",
		&m.Price: "price",
	}
}

func GetProdudcts() (products []model.Product, err error) {
	fmt.Println("Get products service")
	data.DB().Table("products").
		Select(`
			id,
			name,
			price
		`).Order("name ASC").Scan(&products)

	if len(products) == 0 {
		err = errors.New("Unable to find any products!")
	}

	return products, err
}
