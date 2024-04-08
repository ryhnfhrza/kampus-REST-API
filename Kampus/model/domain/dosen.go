package domain

import "time"

type Dosen struct {
	Id           int
	Nama         string
	Gender       string
	TanggalLahir time.Time
}