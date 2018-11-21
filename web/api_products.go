package web

import (
	"fmt"
	"net/http"

	"github.com/xyingsoft/golang-vue/service"
)

func getProducts(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("api")
	// params := mux.Vars(r)

	products, error := service.GetProdudcts()

	if error != nil {
		fmt.Println("ERROR")
	}

	return JsonAction{products}
}

func getProduct(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("get product api")

	return JsonAction{}
}

func createProduct(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("create product api")

	return JsonAction{}
}

func updateProduct(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("update product api")

	return JsonAction{}
}

func deleteProduct(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("delete product api")

	return JsonAction{}
}
