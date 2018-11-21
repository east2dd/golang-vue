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
	Name string `json:"name"`
}

func (m *CategoryParams) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&m.Name: "name",
	}
}

func GetCategories() (categories []model.Category, err error) {
	fmt.Println("Get categories service")
	data.DB().Table("categories").
		Select(`
			id,
			name
		`).Scan(&categories)

	if len(categories) == 0 {
		err = errors.New("Unable to find any categories!")
	}

	return categories, err
}
