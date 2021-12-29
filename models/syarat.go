package models

import (
	// "database/sql"
	"fmt"
	u "go-rest-api/utils"
	"time"
	// "gorm.io/gorm"
)

type Syarat struct {
	ID              int       `form:"id" json:"id" gorm:"primaryKey"`
	UraianKegiatan  string    `form:"uraian_kegiatan" json:"uraian_kegiatan"`
	NamaKegiatan    string    `form:"nama_kegiatan" json:"nama_kegiatan"`
	Kategori        string    `form:"kategori" json:"kategori"`
	TanggalKegiatan time.Time `form:"tanggal_kegiatan" json:"tanggal_kegiatan"`
	RubrikID        int16     `form:"rubrik_id" json:"rubrik_id"`
	PeriodeID       int16     `form:"periode_id" json:"periode_id"`
	Satuan          *float64  `form:"satuan" json:"satuan"`
	CreatedAt       time.Time `form:"created_at" json:"created_at"`
	UpdatedAt       time.Time `form:"updated_at" json:"updated_at"`
	Username        string    `form:"username" json:"username"`
	Dokumens        []Dokumen `gorm:"foreignKey:SyaratsID" json:"dokumen"`
}

func (kegiatan *Syarat) Validate() (map[string]interface{}, bool) {
	if kegiatan.UraianKegiatan == "" {
		return u.Message(false, "uraian kegiatan should be on the payload"), false
	}
	return u.Message(true, "success"), true
}

func (kegiatan *Syarat) Create() map[string]interface{} {
	getDB().Create(kegiatan)
	resp := u.Message(true, "success")
	resp["kegiatan"] = kegiatan
	return resp
}

func GetKegiatan(id int) *Syarat {
	kegiatan := &Syarat{}
	err := getDB().Table("syarats").Preload("Dokumens").Where("id = ?", id).First(kegiatan).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return kegiatan
}

func GetKegiatans(username string) []*Syarat {
	kegiatans := make([]*Syarat, 0)
	err := getDB().Table("syarats").Preload("Dokumens").Where("username = ?", username).Limit(10).Find(&kegiatans).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return kegiatans
}
