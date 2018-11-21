package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xyingsoft/golang-vue/config"
	"github.com/xyingsoft/golang-vue/data"
	"github.com/xyingsoft/golang-vue/web"
)

func main() {
	config.Init()
	data.Init()

	router := mux.NewRouter()
	web.Mount(router)
	fmt.Println("Listening :8000")
	http.ListenAndServe(":8000", router)
}
