package dosenWeb

import "time"

type DosenResponse struct {
	Id           int    `json:"id"`
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
}