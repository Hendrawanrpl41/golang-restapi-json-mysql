package router

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"belajar/RestAPiMYSQL/controller"
	"log"
)

 func Router(){
	fmt.Println("Server Running")
	router := mux.NewRouter().StrictSlash(true)

	//Siswas
	router.HandleFunc("/siswas", controller.InquerySiswas).Methods("GET")
	router.HandleFunc("/siswa", controller.FindSiswa).Methods("GET")
	router.HandleFunc("/siswa", controller.StoreSiswa).Methods("POST")
	router.HandleFunc("/siswa", controller.UpdateSiswa).Methods("PUT")
	router.HandleFunc("/siswa", controller.DeleteSiswa).Methods("DELETE")
	//Kelass
	router.HandleFunc("/kelass", controller.InqueryKelass).Methods("GET")
	router.HandleFunc("/kelas", controller.StoreKelas).Methods("POST")
	router.HandleFunc("/kelas", controller.UpdateKelas).Methods("PUT")
	router.HandleFunc("/kelas", controller.DeleteKelas).Methods("DELETE")

	//server
	log.Fatal(http.ListenAndServe(":8081", router))
 }