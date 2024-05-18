package domain

import "time"

type Mahasiswa struct {
	NIM          string
	Nama         string
	Gender       string
	TanggalLahir time.Time
	Semester     int
	KodeJurusan int
	Angkatan int
	KodeKelas string
}