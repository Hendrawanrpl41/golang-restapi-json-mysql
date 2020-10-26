package model

type Siswa struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	TempatLahir string `json:"tempatLahir"`
	TanggalLahir string `json:"TanggalLahir"`
	Alamat string `json:"alamat"`
	
}

type Kelas struct{
	ID int `json:"id"`
	Kelas string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	Siswa []Siswa
}

type Siswas []Siswa
type Kelass []Kelas


