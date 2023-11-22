package dosenService

import (
	"context"
	"database/sql"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/dosenWeb"
	"kampus/repository/dosenRepository"
)

type DosenServiceImpl struct {
	dosenRepository dosenRepository.DosenRepository
	Db *sql.DB
}

func(dosenService *DosenServiceImpl)Create(ctx context.Context, request dosenWeb.DosenCreateRequest) dosenWeb.DosenResponse{
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen := domain.Dosen{
		Nama: request.Nama,
		Gender: request.Gender,
		Umur: request.Umur,
	}

	dosen = dosenService.dosenRepository.Create(ctx,tx,dosen)

	return helper.ToDosenResponse(dosen)
}

func(dosenService *DosenServiceImpl)Update(ctx context.Context, request dosenWeb.DosenUpdateRequest) dosenWeb.DosenResponse{
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := dosenService.dosenRepository.FindById(ctx,tx,request.Id)
	helper.PanicIfError(err)

	dosen.Nama = request.Nama
	dosen.Umur = request.Umur
	
	dosen = dosenService.dosenRepository.Update(ctx,tx,dosen)

	return helper.ToDosenResponse(dosen)
}

func(dosenService *DosenServiceImpl)Delete(ctx context.Context, dosenId int){
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := dosenService.dosenRepository.FindById(ctx,tx,dosenId)
	helper.PanicIfError(err)
	
	dosenService.dosenRepository.Delete(ctx,tx,dosen)
}

func(dosenService *DosenServiceImpl)FindById(ctx context.Context, dosenId int) dosenWeb.DosenResponse{
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := dosenService.dosenRepository.FindById(ctx,tx,dosenId)
	helper.PanicIfError(err)

	return helper.ToDosenResponse(dosen)
}

func(dosenService *DosenServiceImpl)FindAll(ctx context.Context) []dosenWeb.DosenResponse{
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen := dosenService.dosenRepository.FindAll(ctx,tx)
	helper.PanicIfError(err)

	return helper.ToDosenResponses(dosen)
}
