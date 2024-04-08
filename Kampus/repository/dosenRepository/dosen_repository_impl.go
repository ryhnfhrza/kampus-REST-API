package dosenRepository

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)

type DosenRepositoryImpl struct{

}

func NewDosenRepositoryImpl()DosenRepository{
	return &DosenRepositoryImpl{}
}

func(Repository *DosenRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, dosen domain.Dosen) domain.Dosen{
	SQL := "insert into dosen (id,nama,gender,tanggal_lahir) values (?,?,?,?)"
	result,err := tx.ExecContext(ctx,SQL,dosen.Id,dosen.Nama,dosen.Gender,dosen.TanggalLahir)
	helper.PanicIfError(err)

	id,err := result.LastInsertId()
	helper.PanicIfError(err)

	dosen.Id = int(id)

	return dosen
}

func(Repository *DosenRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, dosen domain.Dosen) domain.Dosen{
	SQL := "update dosen set nama = ? where id = ?"
	_,err := tx.ExecContext(ctx,SQL,dosen.Nama,dosen.Id)
	helper.PanicIfError(err)
	
	return dosen
}

func(Repository *DosenRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, dosen domain.Dosen){
	SQL := "delete from dosen where id = ?"
	_,err := tx.ExecContext(ctx,SQL,dosen.Id)
	helper.PanicIfError(err)
}

func(Repository *DosenRepositoryImpl)FindById(ctx context.Context, tx *sql.Tx,dosenId int)  (domain.Dosen, error){
	SQL := "select id,nama,gender,tanggal_lahir from dosen where id = ?"
	rows,err := tx.QueryContext(ctx,SQL,dosenId)
	helper.PanicIfError(err)
	defer rows.Close()

	dosen := domain.Dosen{}
	
	if rows.Next(){
		err := rows.Scan(&dosen.Id,&dosen.Nama,&dosen.Gender,&dosen.TanggalLahir)
		helper.PanicIfError(err)
		return dosen,nil
	}else{
		return dosen,errors.New("id dosen tidak ditemukan")
	}
}

func(Repository *DosenRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Dosen{
	SQL := "select id,nama,gender,tanggal_lahir from dosen"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var dosen []domain.Dosen

	for rows.Next(){
		dsn := domain.Dosen{}
		err := rows.Scan(&dsn.Id,&dsn.Nama,&dsn.Gender,&dsn.TanggalLahir)
		helper.PanicIfError(err)
		dosen = append(dosen, dsn)
	}
	return dosen
}
