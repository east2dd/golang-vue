package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xyingsoft/golang-vue/app"
	"github.com/xyingsoft/golang-vue/controllers"

	"github.com/gorilla/mux"
)

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func main() {
	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "./public/index.html", "entry point")
	flag.StringVar(&static, "static", "./public", "static files")
	flag.Parse()

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := mux.NewRouter()
	// Handle all preflight request for CORS
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	// api route
	router.Use(app.JwtAuthentication)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/{id}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/api/categories/{id}/products", controllers.GetProductsFor).Methods("GET")
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")

	// static service for vue
	// router.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))
	// router.PathPrefix("/").HandlerFunc(IndexHandler(entry))

	fmt.Println("Listening: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
