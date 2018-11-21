package service

import (
	"errors"
	"fmt"

	"github.com/guregu/null"
	"github.com/xyingsoft/golang-vue/data"
)

type ProductHash struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type ProductParams struct {
	Name  null.String `json:"name"`
	Price int64       `json:"price"`
}

func GetProdudcts() (products []ProductHash, err error) {
	fmt.Println("Get products service")
	data.DB().Table("products").
		Select(`
					name
			`).
		Order("name ASC").
		Scan(&products)

	if len(products) == 0 {
		fmt.Println("Unable to find any products!")
		err = errors.New("Unable to find any products!")
	}

	return products, err
}
