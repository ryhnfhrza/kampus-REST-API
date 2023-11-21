package mahasiswa

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)

type MahasiswaRepositoryImpl struct{

}

func(mahasiswaRepository *MahasiswaRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "insert into mahasiswa (nim,nama,gender,umur,semester) values (?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM,mahasiswa.Nama,mahasiswa.Gender,mahasiswa.Umur,mahasiswa.Semester)
	helper.PanicIfError(err)

	return mahasiswa
}

func(mahasiswaRepository *MahasiswaRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "update mahasiswa set nama = ? ,umur = ?,semester = ? where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.Nama,mahasiswa.Umur,mahasiswa.Semester,mahasiswa.NIM)
	helper.PanicIfError(err)

	return mahasiswa
}

func(mahasiswaRepository *MahasiswaRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa){
	SQL := "delete from mahasiswa where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM)
	helper.PanicIfError(err)
}

func(mahasiswaRepository *MahasiswaRepositoryImpl)FindByNim(ctx context.Context, tx *sql.Tx, mahasiswaNIM string) (domain.Mahasiswa, error){
	SQL := "select nim,nama,gender,umur,semester from mahasiswa where nim = ?"
	rows,err := tx.QueryContext(ctx,SQL,mahasiswaNIM)
	helper.PanicIfError(err)

	defer rows.Close()

	mahasiswa := domain.Mahasiswa{}
	if rows.Next(){
		err := rows.Scan(&mahasiswa.NIM,&mahasiswa.Nama,&mahasiswa.Gender,&mahasiswa.Umur,&mahasiswa.Semester)
		helper.PanicIfError(err)
		return mahasiswa,nil
	}else{
		return mahasiswa,errors.New("NIM tidak ditemukan")
	}
}

func(mahasiswaRepository *MahasiswaRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa{
	SQL := "select nim,nama,gender,umur,semester from mahasiswa"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var mahasiswa []domain.Mahasiswa

	for rows.Next(){
		mhs := domain.Mahasiswa{}
		err := rows.Scan(&mhs.NIM,&mhs.Nama,&mhs.Gender,&mhs.Umur,&mhs.Semester)
		helper.PanicIfError(err)
		mahasiswa = append(mahasiswa, mhs)
	}
	return mahasiswa
}
