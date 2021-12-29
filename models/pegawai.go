package models

import (
	"fmt"
	"time"
)

type Pegawai struct {
	ID uint `json:"id" gorm:"primaryKey"`
	SyaratsID uint `json:"syarats_id"`
	Nama string `json:"nama"`
	Nip string `json:"nip"`
	Posisi *string `json:"posisi"`
	Rubrik_ID int `json:"rubrik_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PegawaiJoin struct{
	Pegawai
	UraianKegiatan string `json:"uraian_kegiatan"`
	NamaKegiatan string `json:"nama_kegiatan"`
	Kategori string `json:"kategori"`
	TanggalKegiatan time.Time `json:"tanggal_kegiatan"`
	RubrikID int `json:"rubrik_id"`
	PeriodeID int `json:"periode_id"`
	Satuan *int `json:"satuan"`
	Username string `json:"username"`
}

type PegawaiDoc struct {
	PegawaiJoin
	Dokumen []Dokumen `json:"dokumen"`
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
	pegawai := make([]*PegawaiJoin, 0)
	data := make([]*PegawaiDoc, 0)

	err := db.Raw("SELECT pegawais.*, syarats.uraian_kegiatan, syarats.nama_kegiatan, syarats.kategori, syarats.tanggal_kegiatan, syarats.rubrik_id, syarats.periode_id, syarats.satuan, syarats.username FROM pegawais JOIN syarats ON pegawais.syarats_id = syarats.id WHERE nip LIKE ?", nip).Scan(&pegawai).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, v := range pegawai {
		var row PegawaiDoc
		row.PegawaiJoin = *v
		err := getDB().Table("dokumens").Where("syarats_id = ?", v.SyaratsID).Find(&row.Dokumen).Error
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, &row)
	}
	return data
}