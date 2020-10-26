package controller

import(
	// "fmt"
	"net/http"
	"encoding/json"
	"belajar/RestAPiMYSQL/dao"
)

func InquerySiswas(res http.ResponseWriter, req *http.Request){
	var siswa dao.Siswas
	result := siswa.FindAll()
	json.NewEncoder(res).Encode(result)
}

func StoreSiswa(res http.ResponseWriter, req *http.Request){
	var siswa dao.Siswa
	json.NewDecoder(req.Body).Decode(&siswa)
	// controller.Store(Siswa)
	// fmt.Println(siswa)
	result := siswa.Create()
	// controller.Save(Siswa)
	json.NewEncoder(res).Encode(result)
}

//findOne
func FindSiswa(res http.ResponseWriter, req *http.Request){
	var siswa dao.Siswa
	json.NewDecoder(req.Body).Decode(&siswa)
	// controller.Store(Siswa)
	result := siswa.FindOne(siswa.ID)
	// controller.Save(Siswa)
	json.NewEncoder(res).Encode(result)
}
//update
func UpdateSiswa(res http.ResponseWriter, req *http.Request){
	var siswa dao.Siswa
	json.NewDecoder(req.Body).Decode(&siswa)
	// controller.Store(Siswa)
	result := siswa.Update()
	// controller.Save(Siswa)
	json.NewEncoder(res).Encode(result)
}

//delete
func DeleteSiswa(res http.ResponseWriter, req *http.Request){
	var siswa dao.Siswa
	json.NewDecoder(req.Body).Decode(&siswa)
	// controller.Store(Siswa)
	result := siswa.Delete()
	// controller.Save(Siswa)
	json.NewEncoder(res).Encode(result)
}