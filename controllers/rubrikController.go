package controllers

import (
	"encoding/json"
	"go-rest-api/utils"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RubrikResp struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []*Rubriks `json:"data"`
}

type RubrikByIDResp struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    *Rubriks `json:"data"`
}

type Rubriks struct {
	ID                 uint                  `json:"id"`
	KategoriRubrikID   uint                  `json:"kategori_rubrik_id"`
	Uraian             string                `json:"uraian"`
	KategoriSatuanID   uint                  `json:"kategori_satuan_id"`
	Syarat             string                `json:"syarat"`
	Keterangan         string                `json:"keterangan"`
	CreatedAt          string                `json:"created_at"`
	UpdatedAt          string                `json:"updated_at"`
	RubrikDetail       []RubrikDetail        `json:"rubrik_detail"`
	Kategori           Kategori              `json:"kategori"`
	KategoriSatuan     KategoriSatuan        `json:"kategori_satuan"`
	SyaratRubrikDetail []*SyaratRubrikDetail `json:"syarat_rubrik_detail"`
}

type RubrikDetail struct {
	ID                      uint      `json:"id"`
	RubrikID                uint      `json:"rubrik_id"`
	PelaksanaID             uint      `json:"pelaksana_id"`
	IsShow                  int       `json:"is_show"`
	PerluVerifikasiKeuangan int       `json:"perlu_verifikasi_keuangan"`
	Nilai                   float32   `json:"nilai"`
	CreatedAt               string    `json:"created_at"`
	UpdatedAt               string    `json:"updated_at"`
	Uraian                  string    `json:"uraian"`
	Pelaksana               Pelaksana `json:"pelaksana"`
	Rubrik                  Rubrik    `json:"rubrik"`
}

type Pelaksana struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Kode      string `json:"kode"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Rubrik struct {
	ID               uint   `json:"id"`
	KategoriRubrikID uint   `json:"kategori_rubrik_id"`
	Uraian           string `json:"uraian"`
	KategoriSatuanID uint   `json:"kategori_satuan_id"`
	Syarat           string `json:"syarat"`
	Keterangan       string `json:"keterangan"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type Kategori struct {
	ID         uint    `json:"id"`
	Nama       string  `json:"nama"`
	Keterangan *string `json:"keterangan"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type KategoriSatuan struct {
	ID         uint    `json:"id"`
	Nama       string  `json:"nama"`
	Status     string  `json:"status"`
	Keterangan *string `json:"keterangan"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type SyaratRubrik struct {
	ID         uint    `json:"id"`
	Nama       string  `json:"nama"`
	Slug       string  `json:"slug"`
	Keterangan *string `json:"keterangan"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type SyaratRubrikDetail struct {
	ID             uint         `json:"id"`
	RubrikID       uint         `json:"rubrik_id"`
	SyaratRubrikID uint         `json:"syarat_rubrik_id"`
	Keterangan     *string      `json:"keterangan"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	SyaratRubrik   SyaratRubrik `json:"syarat_rubrik"`
}

type SyaratResp struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []*SyaratRubrik `json:"data"`
}

var ApiSiremunEndpoint string = os.Getenv("API_SIREMUN")

func GetRubriks() []*Rubriks {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ApiSiremunEndpoint+"rubrik", nil)
	req.Header.Add("Authorization-Nugra", "bnVncmEgZ2FudGVuZyBtYWtzaW1hbA==")
	res, _ := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	var data RubrikResp
	errr := json.NewDecoder(res.Body).Decode(&data)
	if errr != nil {
		log.Fatal(errr)
	}
	return data.Data
}

func GetRubrik(id string) *Rubriks {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ApiSiremunEndpoint+"rubrik/"+id, nil)
	req.Header.Add("Authorization-Nugra", "bnVncmEgZ2FudGVuZyBtYWtzaW1hbA==")
	res, _ := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	var data RubrikByIDResp
	errr := json.NewDecoder(res.Body).Decode(&data)
	if errr != nil {
		log.Fatal(errr)
	}
	return data.Data
}

func GetAllSyarats() []*SyaratRubrik {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ApiSiremunEndpoint+"syarat/get-all", nil)
	req.Header.Add("Authorization-Nugra", "bnVncmEgZ2FudGVuZyBtYWtzaW1hbA==")
	res, _ := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var data *SyaratResp

	errr := json.NewDecoder(res.Body).Decode(&data)
	if errr != nil {
		log.Fatal(errr)
	}
	return data.Data
}

var GetRubrikAll = func(w http.ResponseWriter, r *http.Request) {
	data := GetRubriks()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetRubrikByID = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := GetRubrik(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetSyaratAll = func(w http.ResponseWriter, r *http.Request) {
	data := GetAllSyarats()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
