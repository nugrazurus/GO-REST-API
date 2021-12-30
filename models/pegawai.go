package models

import (
	"fmt"
	"time"
)

type Pegawai struct {
	ID uint `json:"id"`
	SyaratsID uint `json:"syarats_id" gorm:"primaryKey"`
	Nama string `json:"nama"`
	Nip string `json:"nip"`
	Posisi *string `json:"posisi"`
	Rubrik_ID int `json:"rubrik_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Syarat Syarat `json:"syarat" gorm:"foreignKey:SyaratsID"`
}

type PegawaiDoc struct {
	Pegawai
	UraianKegiatan string `json:"uraian_kegiatan"`
	NamaKegiatan string `json:"nama_kegiatan"`
	Kategori string `json:"kategori"`
	TanggalKegiatan time.Time `json:"tanggal_kegiatan"`
	RubrikID int `json:"rubrik_id"`
	PeriodeID int `json:"periode_id"`
	Satuan *int `json:"satuan"`
	Username string `json:"username"`
	Dokumen []Dokumen `json:"dokumen" gorm:"foreignKey:SyaratsID"`
}

func GetPegawai(id int) *Pegawai  {
	pegawai := &Pegawai{}

	err := getDB().Table("pegawais").Where("id = ?", id).First(pegawai).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pegawai
}

func GetPegawais(syarats_id int) []*Pegawai   {
	pegawai := make([]*Pegawai, 0)

	err := getDB().Table("pegawais").Where("syarats_id = ?", syarats_id).Find(&pegawai).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pegawai
}

func GetKegiatanPegawai(nip string) []*PegawaiDoc  {
	data := make([]*PegawaiDoc, 0)

	err := getDB().Table("pegawais").Select("pegawais.*, syarats.uraian_kegiatan, syarats.nama_kegiatan, syarats.kategori, syarats.tanggal_kegiatan, syarats.rubrik_id, syarats.periode_id, syarats.satuan, syarats.username").Joins("LEFT JOIN syarats ON pegawais.syarats_id = syarats.id").Where("nip LIKE ?", nip).Preload("Dokumen", "username LIKE ? OR username LIKE ?", "%op%", nip).Find(&data).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return data
}

func GetKegiatanPegawaiTwo(nip string) []*Pegawai  {
	pegawai := make([]*Pegawai, 0)

	err := getDB().Table("pegawais").Preload("Syarat.Dokumens").Where("nip LIKE ?", nip).Find(&pegawai).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pegawai
}