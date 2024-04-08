package dosenService

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/dosenWeb"
	"kampus/repository/dosenRepository"
	"time"

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

func(Service *DosenServiceImpl)Create(ctx context.Context, request dosenWeb.DosenCreateRequest) dosenWeb.DosenResponse{
	err := Service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	formatter := "2006-01-02"
	birthDate := request.TanggalLahir
	parseTime , err := time.Parse(formatter,birthDate)
	helper.PanicIfError(err)

	dosen := domain.Dosen{
		Nama: request.Nama,
		Gender: request.Gender,
		TanggalLahir: parseTime,
	}

	dosen = Service.dosenRepository.Create(ctx,tx,dosen)

	return helper.ToDosenResponse(dosen)
}

func(Service *DosenServiceImpl)Update(ctx context.Context, request dosenWeb.DosenUpdateRequest) dosenWeb.DosenResponse{
	err := Service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := Service.dosenRepository.FindById(ctx,tx,request.Id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	request.Nama = helper.GetDefaultIfEmpty(request.Nama,dosen.Nama)
	
	dosen.Nama = request.Nama
	
	dosen = Service.dosenRepository.Update(ctx,tx,dosen)

	return helper.ToDosenResponse(dosen)
}

func(Service *DosenServiceImpl)Delete(ctx context.Context, dosenId int){
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := Service.dosenRepository.FindById(ctx,tx,dosenId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	Service.dosenRepository.Delete(ctx,tx,dosen)
}

func(Service *DosenServiceImpl)FindById(ctx context.Context, dosenId int) dosenWeb.DosenResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen,err := Service.dosenRepository.FindById(ctx,tx,dosenId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToDosenResponse(dosen)
}

func(Service *DosenServiceImpl)FindAll(ctx context.Context) []dosenWeb.DosenResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosen := Service.dosenRepository.FindAll(ctx,tx)

	return helper.ToDosenResponses(dosen)
}
