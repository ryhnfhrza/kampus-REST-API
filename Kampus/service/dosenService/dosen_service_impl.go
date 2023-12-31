package dosenService

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/dosenWeb"
	"kampus/repository/dosenRepository"

	"github.com/go-playground/validator/v10"
)

type DosenServiceImpl struct {
	dosenRepository dosenRepository.DosenRepository
	Db *sql.DB
	validate *validator.Validate
}

func NewDosenService(DosenRepository dosenRepository.DosenRepository,db *sql.DB,Validate *validator.Validate)DosenService{
	return &DosenServiceImpl{
		dosenRepository: DosenRepository,
		Db: db,
		validate: Validate,
	}
}

func(dosenService *DosenServiceImpl)Create(ctx context.Context, request dosenWeb.DosenCreateRequest) dosenWeb.DosenResponse{
	err := dosenService.validate.Struct(request)
	helper.PanicIfError(err)

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
	err := dosenService.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := dosenService.dosenRepository.FindById(ctx,tx,request.Id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

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
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	dosenService.dosenRepository.Delete(ctx,tx,dosen)
}

func(dosenService *DosenServiceImpl)FindById(ctx context.Context, dosenId int) dosenWeb.DosenResponse{
	tx,err := dosenService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := dosenService.dosenRepository.FindById(ctx,tx,dosenId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

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
