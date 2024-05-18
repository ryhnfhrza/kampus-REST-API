package matakuliahWeb

type MatakuliahUpdateRequest struct {
	Kode       string `validate:"required,max=20,min=1" json:"kode"`
	Matakuliah string `json:"matakuliah"`
	SKS        int    `validate:"max=15" json:"sks"`
}