package dao

import(
	"fmt"
	"belajar/RestAPiMYSQL/model"
	"belajar/RestAPiMYSQL/config"
)

type Siswa model.Siswa
type Siswas model.Siswas

//findAll
func (siswa Siswas) FindAll() interface{} {
	result := config.Db.Find(&siswa)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(siswa)
	return result.Value
}

//findOne
func (siswa Siswa) FindOne(id int) interface{} {
	result := config.Db.First(&siswa, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(siswa)
	return result.Value
}

//store data
func (siswa Siswa) Create() interface{}{
	result := config.Db.Create(&siswa)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(siswa)
	return result
}

//update
func (siswa Siswa) Update() interface{} {
	result := config.Db.Save(&siswa)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(siswa)
	return result.Value
}

//delete
func (siswa Siswa) Delete() interface{} {
	result := config.Db.Delete(&siswa, siswa.ID)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(siswa)
	return result.Value
}
