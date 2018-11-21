package main

import (
	"flag"
	"fmt"
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

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	flag.StringVar(&entry, "entry", "./public/index.html", "the entrypoint to serve.")
	flag.StringVar(&static, "static", "./public", "the directory to serve static files from.")
	flag.Parse()

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/{id}/products", controllers.GetProductsFor).Methods("GET")
	router.HandleFunc("/api/me/categories", controllers.GetCategoriesFor).Methods("GET")

	router.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))
	router.PathPrefix("/").HandlerFunc(IndexHandler(entry))

	fmt.Println("Listening: " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
