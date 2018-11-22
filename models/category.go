package models

import (
	u "github.com/xyingsoft/golang-vue/utils"
)

type CategoryParams struct {
	Name string
}

type Category struct {
	ID       uint
	Name     string
	Products []*Product
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

	res, err := db.Exec(`INSERT INTO categories(name) VALUES( ? )`, category.Name)

	if err == nil {
		id, err := res.LastInsertId()
		checkErr(err)

		category.ID = uint(id)
	}

	if category.ID <= 0 {
		return u.Message(false, "Failed to create")
	}

	resp := u.Message(true, "success")
	resp["category"] = category
	return resp
}
