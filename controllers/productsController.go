package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xyingsoft/golang-vue/models"
	u "github.com/xyingsoft/golang-vue/utils"
)

var GetProducts = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProducts()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := product.Create()
	u.Respond(w, resp)
}

var UpdateProduct = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	product := models.GetProduct(uint(id))
	product.Name = params["name"]
	product.Update()

	resp := u.Message(true, "success")
	resp["data"] = product
	u.Respond(w, resp)
}
