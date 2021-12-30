package controllers

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateKegiatan = func (w http.ResponseWriter, r *http.Request)  {
	resp := utils.Message(true, "success")

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	file, handler, err := r.FormFile("file")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	fileLocation := filepath.Join(dir, "storage", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := r.PostForm
	resp["data"] = data
	utils.Respond(w, resp)
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