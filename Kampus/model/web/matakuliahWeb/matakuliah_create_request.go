package matakuliahWeb

type MatakuliahCreateRequest struct {
	Kode       string `Validate:"required,max=20,min=1"`
	Matakuliah string `Validate:"required,max=100,min=1"`
	SKS        int    `Validate:"required,max=15,min=1"`
}