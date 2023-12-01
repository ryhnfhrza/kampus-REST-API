package matakuliahWeb

type MatakuliahUpdateRequest struct {
	Kode       string `validate:"required,max=20,min=1" json:"kode"`
	Matakuliah string `validate:"required,max=100,min=1" json:"matakuliah"`
	SKS        int    `validate:"required,max=15,min=1" json:"sks"`
}