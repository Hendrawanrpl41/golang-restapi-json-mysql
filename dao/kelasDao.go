package dao

import(
	"fmt"
	"belajar/RestAPiMYSQL/model"
	"belajar/RestAPiMYSQL/config"
)

type Kelas model.Kelas
type Kelass model.Kelass

//Kelass
func (kelas Kelas) Create() interface{}{
	result := config.Db.Create(&kelas)
	if result.Error != nil {
		return result.Error
	}
	return result
}

//findAll
func (kelas Kelass) FindAll() interface{} {
	fmt.Println()
	result := config.Db.Find(&kelas)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//findOne
func (kelas Kelas) FindOne() interface{} {
	fmt.Println()
	result := config.Db.Find(&kelas)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result)
	return result
}

//update
func (kelas Kelas) Update() interface{} {
	result := config.Db.Save(&kelas)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(kelas)
	return result
}

//delete
func (kelas Kelas) Delete() interface{} {
	result := config.Db.Delete(&kelas, kelas.ID)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(kelas)
	return result
}

