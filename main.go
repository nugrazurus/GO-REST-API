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

	router.HandleFunc("/api/dokumen/create", controllers.CreateDocument).Methods("POST")
	router.HandleFunc("/api/dokumen/{id}", controllers.GetDocument).Methods("GET")
	router.HandleFunc("/api/dokumens/{syarats_id}", controllers.GetDocuments).Methods("GET")
	router.HandleFunc("/api/syarat/{id}", controllers.GetKegiatan).Methods("GET")
	router.HandleFunc("/api/syarats/{username}", controllers.GetKegiatans).Methods("GET")

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