package controllers

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var createKegiatan = func (w http.ResponseWriter, r *http.Request)  {
	
}

var GetKegiatan = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	s, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	data := models.GetKegiatan(s)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetKegiatans = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	data := models.GetKegiatans(vars["username"])
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}