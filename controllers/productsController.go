package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xyingsoft/golang-vue/models"
	u "github.com/xyingsoft/golang-vue/utils"
)

var GetProduct = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "Bad Request"), http.StatusBadRequest)
		return
	}

	data := models.GetProduct(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, http.StatusOK)
}

var GetProducts = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProducts()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, http.StatusOK)
}

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		u.Respond(w, u.Message(false, "Bad Request"), http.StatusBadRequest)
		return
	}

	resp := product.Create()
	u.Respond(w, resp, http.StatusOK)
}

var UpdateProduct = func(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value("user").(uint)
	// if user < 1 {
	// 	u.Respond(w, u.Message(false, "Bad Request"), http.StatusBadRequest)
	// 	return
	// }

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Bad Request"), http.StatusBadRequest)
		return
	}

	decorder := json.NewDecoder(r.Body)
	productParams := &models.ProductParams{}
	decorder.Decode(&productParams)

	product := models.GetProduct(uint(id))
	product.Name = productParams.Name
	product.Description = productParams.Description
	product.Update()

	resp := u.Message(true, "success")
	resp["data"] = product
	u.Respond(w, resp, http.StatusOK)
}
