package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getSyaratAll(w http.ResponseWriter, r *http.Request) {
	var syarats Syarats
	var arr_syarat []Syarats
	var response ResponseSyarat

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from syarats ORDER BY id DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for rows.Next() {
		if err := rows.Scan(&syarats.Id,
			&syarats.Uraian_kegiatan,
			&syarats.Nama_kegiatan,
			&syarats.Kategori,
			&syarats.Tanggal_kegiatan,
			&syarats.Rubrik_id,
			&syarats.Periode_id,
			&syarats.Created_at,
			&syarats.Updated_at,
			&syarats.Username,
			); err != nil {
			fmt.Println(err)
		} else {
			arr_syarat = append(arr_syarat, syarats)
		}
		count++
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_syarat
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)	
}

func getDokumenAll(w http.ResponseWriter, r *http.Request) {
	var dokumens Dokumens
	var arr_dokumen []Dokumens
	var response ResponseDokumens

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dokumens ORDER BY id DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&dokumens.Id,
			&dokumens.Syarats_id,
			&dokumens.File,
			&dokumens.Nama_asli,
			&dokumens.Tipe_dokumen,
			&dokumens.Created_at,
			&dokumens.Updated_at,
			); err != nil {
				fmt.Println(err)
			} else {
			arr_dokumen = append(arr_dokumen, dokumens)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_dokumen
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)	
}

func getUnitKerjaAll() {
	var unitkerjas UnitKerjas;
	var arr_unitkerjas []UnitKerjas;
	db := connect()
	defer db.Close()
	rows, err := db.Query("Select id,nama,kategori from unit_kerjas ORDER BY id DESC")
	if err != nil {
		log.Print(err)
	}
	for rows.Next(){
		if err := rows.Scan(&unitkerjas.Id, &unitkerjas.Nama, &unitkerjas.Kategori); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_unitkerjas = append(arr_unitkerjas, unitkerjas)
			// fmt.Println(arr_unitkerjas[count])
		}
	}
}

func getKegiatanByNip(w http.ResponseWriter, r *http.Request)  {
	var kegiatans Kegiatans
	var documents Documents
	var arr_kegiatan []Kegiatans
	var response ResponseKegiatans
	vars := mux.Vars(r)
	nip := vars["nip"]
	db := connect()
	defer db.Close()
	
	rows, err := db.Query("SELECT syarats_id AS id, syarats.uraian_kegiatan, syarats.kategori, syarats.nama_kegiatan, syarats.rubrik_id, syarats.periode_id, syarats.tanggal_kegiatan, posisi  FROM pegawais INNER JOIN syarats ON pegawais.syarats_id = syarats.id WHERE pegawais.nip = ? ORDER BY id DESC", nip)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var arr_document []Documents
		if err := rows.Scan(
			&kegiatans.Id,
			&kegiatans.Uraian,
			&kegiatans.Kategori,
			&kegiatans.Nama_kegiatan,
			&kegiatans.Rubrik_id,
			&kegiatans.Periode_id,
			&kegiatans.Tanggal_kegiatan,
			&kegiatans.Posisi,
		); err != nil {
			log.Fatal(err)
		} else {
			rows, err := db.Query("SELECT id,nama_asli,tipe_dokumen FROM dokumens WHERE syarats_id = ?", kegiatans.Id)
			if err != nil {
				log.Fatal(err)
			}
			for rows.Next() {
				if err := rows.Scan(
					&documents.Id,
					&documents.Nama_dokumen,
					&documents.Tipe_dokumen,
				); err != nil {
					log.Fatal(err)
				} else {
					documents.Tipe_id = "1"
					arr_document = append(arr_document, documents)
				}
			}
			kegiatans.Syarat = arr_document
			arr_kegiatan = append(arr_kegiatan, kegiatans)
		}
	}
	response.Status = true
	response.Message = "Success"
	response.Data = arr_kegiatan
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}