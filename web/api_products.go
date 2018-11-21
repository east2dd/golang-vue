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

func getProducts(c http.ResponseWriter, r *http.Request) RenderAction {

	products, error := service.GetProdudcts()

	if error != nil {
		fmt.Println("ERROR")
	}

	return JsonAction{products}
}

func getProduct(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	product := model.Product{}

	data.DB().First(&product, id)

	if product.ID > 0 {
		return JsonAction{product}
	} else {
		return NotFoundAction{}
	}
}

func createProduct(c http.ResponseWriter, r *http.Request) RenderAction {
	params := new(service.ProductParams)
	errs := binding.Bind(r, params)

	if errs != nil {
		return JsonAction{}
	}

	product := &model.Product{Name: params.Name, Price: params.Price}
	data.DB().Create(product)

	return JsonAction{product}
}

func updateProduct(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	params := new(service.ProductParams)

	errs := binding.Bind(r, params)
	if errs != nil {
		return ErrorAction{}
	}

	product := model.Product{}
	data.DB().First(&product, id)
	if product.ID > 0 {
		data.DB().Model(&product).Update(
			&model.Product{
				Name:  params.Name,
				Price: params.Price,
			})
	} else {
		return NotFoundAction{}
	}

	return JsonAction{product}
}

func deleteProduct(c http.ResponseWriter, r *http.Request) RenderAction {
	id := mux.Vars(r)["id"]
	product := model.Product{}

	data.DB().First(&product, id)
	if product.ID > 0 {
		data.DB().Delete(&product)
	} else {
		return NotFoundAction{}
	}

	return JsonAction{}
}
