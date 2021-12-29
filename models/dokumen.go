package models

import (
	"fmt"
	u "go-rest-api/utils"
	"time"
	// "gorm.io/gorm"
)

type Dokumen struct {
	// gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	SyaratsID   uint   `json:"syarats_id"`
	File        string `json:"file"`
	NamaAsli    string `json:"nama_asli"`
	TipeDokumen string `json:"tipe_dokumen"`
	Username    string    `json:"username"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func (document *Dokumen) Validate() (map[string]interface{}, bool) {
	if document.NamaAsli == "" {
		return u.Message(false, "nama dokumen should be on the payload"), false
	}
	if document.TipeDokumen == "" {
		return u.Message(false, "tipe dokumen should be on the payload"), false
	}
	return u.Message(true, "success"), true
}

func (document *Dokumen) Create() map[string]interface{} {
	if resp, ok := document.Validate(); !ok {
		return resp
	}
	getDB().Create(document)

	resp := u.Message(true, "success")
	resp["document"] = document
	return resp
}

func GetDocument(id int) *Dokumen {
	document := &Dokumen{}
	err := getDB().Table("dokumens").Where("id = ?", id).First(document).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return document
}

func GetDocuments(syarats_id int) []*Dokumen {
	documents := make([]*Dokumen, 0)
	err := getDB().Table("dokumens").Where("syarats_id = ?", syarats_id).Find(&documents).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return documents
}
