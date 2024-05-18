package repository

import (
	"context"
	"database/sql"
	"kampus/helper"
	"kampus/model/domain"
)

type DosenKelasMatkulRepositoryImpl struct{

}

func NewDosenKelasMatkulRepositoryImpl()DosenKelasMatkulRepository{
	return &DosenKelasMatkulRepositoryImpl{}
}

func(Repository *DosenKelasMatkulRepositoryImpl)AjarMatkul(ctx context.Context, tx *sql.Tx, ajarMatkul domain.DosenKelasMatkul) domain.DosenKelasMatkul{
	SQL := "insert into dosen_kelas_matakuliah (id_dosen,kode_matakuliah,kode_kelas) values (?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,ajarMatkul.IdDosen,ajarMatkul.KodeMatkul,ajarMatkul.KodeKelas)
	helper.PanicIfError(err)

	return ajarMatkul
}