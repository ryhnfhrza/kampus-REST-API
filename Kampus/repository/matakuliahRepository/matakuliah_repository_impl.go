package matakuliahRepository

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)


type MatakuliahRepositoryImpl struct{
}

func NewMatakuliahRepositoryImpl()MatakuliahRepository{
	return &MatakuliahRepositoryImpl{}
}

func(matakuliahRepository *MatakuliahRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah{
	SQL := "insert into matakuliah (kode,matakuliah,sks) values (?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,matakuliah.Kode,matakuliah.Mata_kuliah,matakuliah.SKS)
	helper.PanicIfError(err)

	return matakuliah
}

func(matakuliahRepository *MatakuliahRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah{
	SQL := "update matakuliah set matakuliah = ?,sks = ? where kode = ?"
	_,err := tx.ExecContext(ctx,SQL,matakuliah.Mata_kuliah,matakuliah.SKS,matakuliah.Kode)
	helper.PanicIfError(err)

	return matakuliah
}

func(matakuliahRepository *MatakuliahRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah){
	SQL := "delete from matakuliah where kode = ?"
	_,err := tx.ExecContext(ctx,SQL,matakuliah.Kode)
	helper.PanicIfError(err)
}

func(matakuliahRepository *MatakuliahRepositoryImpl)FindByKode(ctx context.Context, tx *sql.Tx,kodeMatakuliah string) (domain.Matakuliah, error){
	SQL := "select kode,matakuliah,sks from matakuliah where kode = ?"
	rows,err := tx.QueryContext(ctx,SQL,kodeMatakuliah)
	helper.PanicIfError(err)
	defer rows.Close()

	matakuliah := domain.Matakuliah{}

	if rows.Next(){
		err := rows.Scan(&matakuliah.Kode,&matakuliah.Mata_kuliah,&matakuliah.SKS)
		helper.PanicIfError(err)
		return matakuliah,nil
	}else{
		return matakuliah,errors.New("NIM tidak ditemukan")
	}
}

func(matakuliahRepository *MatakuliahRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Matakuliah{
	SQL := "select kode,matakuliah,sks from matakuliah "
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var matakuliah []domain.Matakuliah

	for rows.Next(){
		mk := domain.Matakuliah{}
		err := rows.Scan(&mk.Kode,&mk.Mata_kuliah,&mk.SKS)
		helper.PanicIfError(err)
		matakuliah = append(matakuliah, mk)
	}
	return matakuliah
}
