package model

type Format struct {
	Angka1 float64 `json:"angka_1,string" binding:"required"`
	Angka2 float64 `json:"angka_2,string" binding:"required"`
	Deskripsi Deskripsi
}

type Deskripsi struct {
	Desc string `json:"desc" binding:"required"`
}

type Hasil struct {
	Pertambahan float64
	Pengurangan float64
	Perkalian float64
	Pembagian float64
}