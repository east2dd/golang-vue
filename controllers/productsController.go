package controllers

import (
	"net/http"

	"github.com/xyingsoft/golang-vue/models"
	u "github.com/xyingsoft/golang-vue/utils"
)

var GetProducts = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProducts()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
