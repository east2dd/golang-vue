package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/xyingsoft/golang-vue/data"
	"github.com/xyingsoft/golang-vue/model"
	"github.com/xyingsoft/golang-vue/service"
)

func getCategories(c http.ResponseWriter, r *http.Request) RenderAction {

	categorys, error := service.GetProdudcts()

	if error != nil {
		fmt.Println("ERROR")
	}

	return JsonAction{categorys}
}

func getCategory(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	category := model.Category{}

	data.DB().First(&category, id)

	if category.ID > 0 {
		return JsonAction{category}
	} else {
		return NotFoundAction{}
	}
}

func createCategory(c http.ResponseWriter, r *http.Request) RenderAction {
	params := new(service.categoryParams)
	errs := binding.Bind(r, params)

	if errs != nil {
		return JsonAction{}
	}

	category := &model.Category{Name: params.Name}
	data.DB().Create(category)

	return JsonAction{category}
}

func updateCategory(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	params := new(service.categoryParams)

	errs := binding.Bind(r, params)
	if errs != nil {
		return ErrorAction{}
	}

	category := model.Category{}
	data.DB().First(&category, id)
	if category.ID > 0 {
		data.DB().Model(&category).Update(
			&model.Category{
				Name: params.Name,
			})
	} else {
		return NotFoundAction{}
	}

	return JsonAction{category}
}

func deletecategory(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	category := model.Category{}

	data.DB().First(&category, id)
	if category.ID > 0 {
		data.DB().Delete(&category)
	} else {
		return NotFoundAction{}
	}

	return JsonAction{}
}
