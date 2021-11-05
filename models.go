package main

import (
	"time"
)

type UnitKerjas struct {
	Id       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Kategori string `form:"kategori" json:"kategori"`
}

type Dokumens struct {
	Id           int          `form:"id" json:"id"`
	Syarats_id   int          `form:"syarats_id" json:"syarats_id"`
	File         string       `form:"file" json:"file"`
	Nama_asli    string       `form:"nama_asli" json:"nama_asli"`
	Tipe_dokumen string       `form:"tipe_dokumen" json:"tipe_dokumen"`
	Created_at   time.Time `form:"created_at" json:"created_at"`
	Updated_at   time.Time    `form:"updated_at" json:"updated_at"`
}

type Syarats struct {
	Id               int       `form:"id" json:"id"`
	Uraian_kegiatan  string    `form:"uraian_kegiatan" json:"uraian_kegiatan"`
	Nama_kegiatan    string    `form:"nama_kegiatan" json:"nama_kegiatan"`
	Kategori         string    `form:"kategori" json:"kategori"`
	Tanggal_kegiatan time.Time `form:"tanggal_kegiatan" json:"tanggal_kegiatan"`
	Rubrik_id        int16     `form:"rubrik_id" json:"rubrik_id"`
	Periode_id       int16     `form:"periode_id" json:"periode_id"`
	Created_at       time.Time `form:"created_at" json:"created_at"`
	Updated_at       time.Time `form:"updated_at" json:"updated_at"`
	Username         string    `form:"username" json:"username"`
}

type ResponseSyarat struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Syarats `json:"data"`
}

type ResponseDokumens struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Dokumens `json:"data"`
}

type ResponseUnitKerjas struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []UnitKerjas `json:"data"`
}

type Kegiatans struct {
	Id int `json:"id"`
	Uraian string `json:"uraian"`
	Kategori string `json:"kategori"`
	Nama_kegiatan string `json:"nama_kegiatan"`
	Rubrik_id string `json:"rubrik_id"`
	Periode_id string `json:"periode_id"`
	Tanggal_kegiatan time.Time `json:"tanggal_kegiatan"`
	Posisi string `json:"posisi"`
	Syarat []Documents `json:"syarat"`
	// Cart_syarat []CartSyarat
}

type Documents struct {
	Id int `json:"id"`
	Tipe_dokumen string `json:"tipe_dokumen"`
	Tipe_id string `json:"tipe_id"`
	Nama_dokumen string `json:"nama_dokumen"`
}

type CartSyarat struct {

}

type ResponseKegiatans struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data []Kegiatans `json:"data"`
}