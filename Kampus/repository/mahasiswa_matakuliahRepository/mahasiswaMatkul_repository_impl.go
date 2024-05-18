package repository

import (
	"context"
	"database/sql"
	"kampus/helper"
	"kampus/model/domain"
)

type MahasiswaMatkulRepositoryImpl struct{

}

func NewMahasiswaMatkulRepositoryImpl()MahasiswaMatkulRepository{
	return &MahasiswaMatkulRepositoryImpl{}
}

func(Repository *MahasiswaMatkulRepositoryImpl)AmbilMatkul(ctx context.Context, tx *sql.Tx, ambilMatkul domain.MahasiswaMatkul) domain.MahasiswaMatkul{
	SQL := "insert into mahasiswa_matakuliah (nim,kode_matakuliah) values (?,?)"
	_,err := tx.ExecContext(ctx,SQL,ambilMatkul.NIM,ambilMatkul.KodeMatkul)
	helper.PanicIfError(err)

	return ambilMatkul
}