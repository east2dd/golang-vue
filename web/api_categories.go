package web

import (
	"fmt"
	"net/http"
)

func getCategories(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("get all categories api")

	return JsonAction{}
}

func getCategory(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("get category api")

	return JsonAction{}
}

func createCategory(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("create category api")

	return JsonAction{}
}

func updateCategory(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("update category api")

	return JsonAction{}
}

func deleteCategory(c http.ResponseWriter, r *http.Request) RenderAction {

	fmt.Println("delete category api")

	return JsonAction{}
}
