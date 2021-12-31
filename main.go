package main

import (
	"fmt"
	"go-rest-api/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main()  {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	dokumen := api.PathPrefix("/dokumen").Subrouter()
	syarat := api.PathPrefix("/syarat").Subrouter()
	pegawai := api.PathPrefix("/pegawai").Subrouter()
	rubrik := api.PathPrefix("/rubrik").Subrouter()

	dokumen.HandleFunc("/create", controllers.CreateDocument).Methods("POST")
	dokumen.HandleFunc("/{id}", controllers.GetDocument).Methods("GET")
	dokumen.HandleFunc("/syarat/{syarats_id}", controllers.GetDocuments).Methods("GET")

	syarat.HandleFunc("/", controllers.CreateKegiatan).Methods("POST")
	syarat.HandleFunc("/{id}", controllers.GetKegiatan).Methods("GET")
	syarat.HandleFunc("/operator/{username}", controllers.GetKegiatans).Methods("GET")
	
	pegawai.HandleFunc("/{id}", controllers.GetPegawai).Methods("GET")
	pegawai.HandleFunc("/syarat/{syarats_id}", controllers.GetPegawais).Methods("GET")
	pegawai.HandleFunc("/kegiatanbynip/{nip}", controllers.GetKegiatanPegawai).Methods("GET")
	
	rubrik.HandleFunc("/all", controllers.GetRubrikAll).Methods("GET")
	rubrik.HandleFunc("/{id}", controllers.GetRubrikByID).Methods("GET")

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}