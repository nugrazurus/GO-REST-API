package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	u "go-rest-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateDocument = func (w http.ResponseWriter, r *http.Request)  {
	document := &models.Dokumen{}

	err := json.NewDecoder(r.Body).Decode(document)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding body"))
		return
	}
	resp := document.Create()
	u.Respond(w, resp)
}

var GetDocument = func (w http.ResponseWriter, r *http.Request)  {
	// id := r.Context().Value("id").(uint)
	vars := mux.Vars(r)
	s, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	data := models.GetDocument(s)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetDocuments = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	s, err := strconv.Atoi(vars["syarats_id"])
	resp := u.Message(true, "success")
	if err != nil {
		fmt.Println(err)
		resp["status"] =  false
		resp["message"] =  err
	}
	data := models.GetDocuments(s)
	resp["data"] = data
	u.Respond(w, resp)
}