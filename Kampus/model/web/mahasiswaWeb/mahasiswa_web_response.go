package mahasiswaWeb

import "time"

type MahasiswaResponse struct {
	NIM          string `json:"nim"`
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	TanggalLahir time.Time   `json:"tanggal_lahir"`
	Semester     int    `json:"semester"`
	KodeJurusan int `json:"kode_jurusan"`
	Angkatan int `json:"angkatan"`
}