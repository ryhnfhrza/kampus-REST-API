package jurusanRepository

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)

type JurusanRepositoryImpl struct{

}

func NewJurusanRepositoryImpl()JurusanRepository{
	return &JurusanRepositoryImpl{}
}

func(repository *JurusanRepositoryImpl)AddJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan) domain.Jurusan{
	SQL := "insert into jurusan (kode,nama) values (?,?)"
	_,err := tx.ExecContext(ctx,SQL,jurusan.Kode,jurusan.NamaJurusan)
	helper.PanicIfError(err)

	return jurusan
}

func(repository *JurusanRepositoryImpl)UpdateNamaJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan) domain.Jurusan{
	SQL := "update jurusan set nama = ? where kode = ?"
	_,err := tx.ExecContext(ctx,SQL,jurusan.NamaJurusan,jurusan.Kode)
	helper.PanicIfError(err)

	return jurusan
}

func(repository *JurusanRepositoryImpl)DeleteJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan){
	SQL := "delete from jurusan where kode = ?"
	_,err := tx.ExecContext(ctx,SQL,jurusan.Kode)
	helper.PanicIfError(err)
}

func(repository *JurusanRepositoryImpl)FindJurusanByKode(ctx context.Context, tx *sql.Tx,jurusanKode int)  (domain.Jurusan, error){
	SQL := "select kode,nama from jurusan where kode = ?"
	rows,err := tx.QueryContext(ctx,SQL,jurusanKode)
	helper.PanicIfError(err)
	
	defer rows.Close()
	
	jurusan :=  domain.Jurusan{}
	if rows.Next(){
		err :=  rows.Scan(&jurusan.Kode,&jurusan.NamaJurusan)
		helper.PanicIfError(err)
		return jurusan,nil
	}else{
		return jurusan,errors.New("Jurusan tidak ditemukan")
	}
}

func(repository *JurusanRepositoryImpl)FindAllJurusan(ctx context.Context, tx *sql.Tx) []domain.Jurusan{
	SQL := "select kode,nama from jurusan "
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	
	defer rows.Close()
	
	var jurusan []domain.Jurusan
	for rows.Next(){
		jrs := domain.Jurusan{}
		err :=  rows.Scan(&jrs.Kode,&jrs.NamaJurusan)
		helper.PanicIfError(err)
		jurusan = append(jurusan, jrs)
	}
	return jurusan
}
