package controller

import(
	"net/http"
	"encoding/json"
	"belajar/RestAPiMYSQL/dao"
)

//Kelas
func InqueryKelass(res http.ResponseWriter, req *http.Request){
	var kelas dao.Kelass
	result := kelas.FindAll()
	json.NewEncoder(res).Encode(result)
}

func StoreKelas(res http.ResponseWriter, req *http.Request){
	var kelas dao.Kelas
	json.NewDecoder(req.Body).Decode(&kelas)
	result := kelas.Create()
	json.NewEncoder(res).Encode(result)
}

//update
func UpdateKelas(res http.ResponseWriter, req *http.Request){
	var kelas dao.Kelas
	json.NewDecoder(req.Body).Decode(&kelas)
	result := kelas.Update()
	json.NewEncoder(res).Encode(result)
}

//delete
func DeleteKelas(res http.ResponseWriter, req *http.Request){
	var kelas dao.Kelas
	json.NewDecoder(req.Body).Decode(&kelas)
	result := kelas.Delete()
	json.NewEncoder(res).Encode(result)
}