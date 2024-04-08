package mahasiswaRepository

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)

type MahasiswaRepositoryImpl struct{

}

func NewMahasiswaRepositoryImpl()MahasiswaRepository{
	return &MahasiswaRepositoryImpl{}
}

func(Repository *MahasiswaRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "insert into mahasiswa (nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan) values (?,?,?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM,mahasiswa.Nama,mahasiswa.Gender,mahasiswa.TanggalLahir,mahasiswa.Semester,mahasiswa.KodeJurusan,mahasiswa.Angkatan)
	helper.PanicIfError(err)

	return mahasiswa
}

func(Repository *MahasiswaRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "update mahasiswa set nama = ? ,tanggal_lahir = ?,semester = ? where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.Nama,mahasiswa.TanggalLahir,mahasiswa.Semester,mahasiswa.NIM)
	helper.PanicIfError(err)

	return mahasiswa
}

func(Repository *MahasiswaRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa){
	SQL := "delete from mahasiswa where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM)
	helper.PanicIfError(err)
}

func(Repository *MahasiswaRepositoryImpl)FindByNim(ctx context.Context, tx *sql.Tx, mahasiswaNIM string) (domain.Mahasiswa, error){
	SQL := "select nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan from mahasiswa where nim = ?"
	rows,err := tx.QueryContext(ctx,SQL,mahasiswaNIM)
	helper.PanicIfError(err)

	defer rows.Close()

	mahasiswa := domain.Mahasiswa{}
	if rows.Next(){
		err := rows.Scan(&mahasiswa.NIM,&mahasiswa.Nama,&mahasiswa.Gender,&mahasiswa.TanggalLahir,&mahasiswa.Semester,&mahasiswa.KodeJurusan,&mahasiswa.Angkatan)
		helper.PanicIfError(err)
		return mahasiswa,nil
	}else{
		return mahasiswa,errors.New("NIM tidak ditemukan")
	}
}

func(Repository *MahasiswaRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa{
	SQL := "select nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan from mahasiswa"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var mahasiswa []domain.Mahasiswa

	for rows.Next(){
		mhs := domain.Mahasiswa{}
		err := rows.Scan(&mhs.NIM,&mhs.Nama,&mhs.Gender,&mhs.TanggalLahir,&mhs.Semester,&mhs.KodeJurusan,&mhs.Angkatan)
		helper.PanicIfError(err)
		mahasiswa = append(mahasiswa, mhs)
	}
	return mahasiswa
}

func(repository *MahasiswaRepositoryImpl)NIMTerakhirMahasiswaPadaJurusan(ctx context.Context,tx *sql.Tx,kodeJurusan int , angkatan int)string{
	SQL := "SELECT nim FROM mahasiswa WHERE kode_jurusan = ? AND angkatan = ? ORDER BY nim DESC LIMIT 1;"
	rows , err := tx.QueryContext(ctx,SQL,kodeJurusan,angkatan)
	helper.PanicIfError(err)
	defer rows.Close()

	var LastInsertNim string
	if rows.Next(){
		err := rows.Scan(&LastInsertNim)
		helper.PanicIfError(err)
		return LastInsertNim
	}else{
		return ""
	}

}
