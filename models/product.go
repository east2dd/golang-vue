package models

import (
	u "github.com/xyingsoft/golang-vue/utils"
)

type ProductParams struct {
	Name string
}

type Product struct {
	ID         uint
	Name       string
	Categories []*Category
}

func (product *Product) Validate() (map[string]interface{}, bool) {
	if product.Name == "" {
		return u.Message(false, "should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (product *Product) Create() map[string]interface{} {
	db := GetDB()
	res, err := db.Exec(`INSERT INTO products(name) VALUES( ? )`, product.Name)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	product.ID = uint(id)
	resp := u.Message(true, "success")
	resp["product"] = product
	return resp
}

func (product *Product) Update() map[string]interface{} {
	res, err := db.Exec(`UPDATE products SET name = ? WHERE id = ?`, product.Name, product.ID)
	checkErr(err)

	var count int64
	count, err = res.RowsAffected()
	checkErr(err)

	if count > 0 {
		resp := u.Message(true, "success")
		resp["product"] = product
		return resp
	} else {
		resp := u.Message(true, "failed")
		return resp
	}
}
