package web

import (
	"github.com/gorilla/mux"
)

func Mount(router *mux.Router) {

	router.HandleFunc("/api/v1/products", HandleAction(getProducts)).Methods("GET")
	router.HandleFunc("/api/v1/products", HandleAction(createProduct)).Methods("POST")
	router.HandleFunc("/api/v1/products/{id}", HandleAction(getProduct)).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(getCategory)).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(deleteCategory)).Methods("DELETE")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(updateCategory)).Methods("UPDATE")

	router.HandleFunc("/api/v1/categories", HandleAction(getCategories)).Methods("GET")
	router.HandleFunc("/api/v1/categories", HandleAction(createCategory)).Methods("POST")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(getCategory)).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(deleteCategory)).Methods("DELETE")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(updateCategory)).Methods("UPDATE")
}
