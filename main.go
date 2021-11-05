package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getsyarat", getSyaratAll).Methods("GET")
	router.HandleFunc("/getdokumen", getDokumenAll).Methods("GET")
	router.HandleFunc("/getkegiatan/{nip}", getKegiatanByNip).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

