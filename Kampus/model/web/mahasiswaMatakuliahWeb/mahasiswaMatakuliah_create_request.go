package web

type MahasiswaMatkulCreateRequest struct {
	NIM        string   `validate:"required,max=15,min=1" json:"nim"`
	KodeMatkul []string `validate:"required,dive,max=20,min=1" json:"kode_matkul"`
}