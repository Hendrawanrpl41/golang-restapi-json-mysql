package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"belajar/RestAPiMYSQL/model"
)

var Db, Err = gorm.Open("mysql", "root:@/sekolah")

func init(){
	if Err != nil{
		panic("Fail Connection Database")
	}

	//model declaration
	var siswa model.Siswa
	var kelas model.Kelas

	//migrate
	Db.AutoMigrate(&siswa)
	Db.AutoMigrate(&kelas)
}