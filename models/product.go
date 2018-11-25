package models

import (
	u "github.com/xyingsoft/golang-vue/utils"
)

type ProductParams struct {
	Name 				string
	Description string
}

type Product struct {
	ID         	uint
	Name       	string
	Description string
	Categories []*Category
}

func (product *Product) Validate() (map[string]interface{}, bool) {
	if product.Name == "" {
		return u.Message(false, "should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (product *Product) Create() map[string]interface{} {
	if resp, ok := product.Validate(); !ok {
		return resp
	}

	db := GetDB()
	res, err := db.Exec(`INSERT INTO products(name, description) VALUES( ?, ? )`, product.Name, product.Description)
	checkErr(err)

	id, res_err := res.LastInsertId()
	checkErr(res_err)

	product.ID = uint(id)
	resp := u.Message(true, "success")
	resp["data"] = product

	return resp
}

func (product *Product) Update() map[string]interface{} {
	res, err := db.Exec(`UPDATE products SET name = ?, description = ? WHERE id = ?`, product.Name, product.Description, product.ID)
	checkErr(err)

	var count int64
	count, err = res.RowsAffected()
	checkErr(err)

	if count > 0 {
		resp := u.Message(true, "success")
		resp["data"] = product
		return resp
	} else {
		resp := u.Message(false, "failed")
		return resp
	}
}
