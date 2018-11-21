package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Mount(router *mux.Router) {

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./views"))))

	router.HandleFunc("/api/v1/products", HandleAction(getProducts)).Methods("GET")
	router.HandleFunc("/api/v1/products", HandleAction(createProduct)).Methods("POST")
	router.HandleFunc("/api/v1/products/{id}", HandleAction(getProduct)).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", HandleAction(deleteProduct)).Methods("DELETE")
	router.HandleFunc("/api/v1/products/{id}", HandleAction(updateProduct)).Methods("PUT")

	router.HandleFunc("/api/v1/categories", HandleAction(getCategories)).Methods("GET")
	router.HandleFunc("/api/v1/categories", HandleAction(createCategory)).Methods("POST")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(getCategory)).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(deleteCategory)).Methods("DELETE")
	router.HandleFunc("/api/v1/categories/{id}", HandleAction(updateCategory)).Methods("PUT")
}
