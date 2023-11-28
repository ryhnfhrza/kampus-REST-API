package matakuliahWeb

type MatakuliahResponse struct {
	Kode       string `json:"kode"`
	Matakuliah string `json:"matakuliah"`
	SKS        int    `json:"sks"`
}