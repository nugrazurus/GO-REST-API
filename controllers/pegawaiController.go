package controllers

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var GetPegawai = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	data := models.GetPegawai(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetPegawais = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	syarats_id, err := strconv.Atoi(vars["syarats_id"])
	if err != nil {
		fmt.Println(err)
	}
	data := models.GetPegawais(syarats_id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetKegiatanPegawai = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	data := models.GetKegiatanPegawai(vars["nip"])
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

