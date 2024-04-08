package jurusanService

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/jurusanWeb"
	jurusanRepository "kampus/repository/jurusanRepository"

	"github.com/go-playground/validator/v10"
)

type JurusanServiceImpl struct{
	JurusanRepository jurusanRepository.JurusanRepository
	Db *sql.DB
	Validate *validator.Validate
}

func NewJurusanService(jurusanRepository jurusanRepository.JurusanRepository,db *sql.DB,validate *validator.Validate)JurusanService{
	return &JurusanServiceImpl{
		JurusanRepository: jurusanRepository,
		Db: db,
		Validate: validate,
	}
}

func(service *JurusanServiceImpl)AddJurusan(ctx context.Context, request jurusanWeb.JurusanCreateRequest) jurusanWeb.JurusanWebResponse{
 err := service.Validate.Struct(request)
 helper.PanicIfError(err)

 tx,err := service.Db.Begin()
 helper.PanicIfError(err)
 defer helper.CommitOrRollback(tx)

 jurusan := domain.Jurusan{
	Kode: request.KodeJurusan,
	NamaJurusan: request.NamaJurusan,
 }

 jurusan = service.JurusanRepository.AddJurusan(ctx,tx,jurusan)

 return helper.ToJurusanResponse(jurusan)
}

func(service *JurusanServiceImpl)UpdateNamaJurusan(ctx context.Context, request jurusanWeb.JurusanUpdateRequest) jurusanWeb.JurusanWebResponse{
 err := service.Validate.Struct(request)
 helper.PanicIfError(err)

 tx,err := service.Db.Begin()
 helper.PanicIfError(err)
 defer helper.CommitOrRollback(tx)

 jurusan,err := service.JurusanRepository.FindJurusanByKode(ctx,tx,request.KodeJurusan)
 if err != nil{
	panic(exception.NewNotFoundError(err.Error()))
 }
 
 jurusan.Kode = request.KodeJurusan
 jurusan.NamaJurusan = request.NamaJurusan

 jurusan = service.JurusanRepository.UpdateNamaJurusan(ctx,tx, jurusan)

 return helper.ToJurusanResponse(jurusan)
}

func(service *JurusanServiceImpl)DeleteJurusan(ctx context.Context, jurusanKode int){
 tx,err := service.Db.Begin()
 helper.PanicIfError(err)
 defer helper.CommitOrRollback(tx)

 jurusan,err := service.JurusanRepository.FindJurusanByKode(ctx,tx,jurusanKode)
 if err != nil{
	panic(exception.NewNotFoundError(err.Error()))
 }
 
 service.JurusanRepository.DeleteJurusan(ctx,tx,jurusan)
}

func(service *JurusanServiceImpl)FindJurusanByKode(ctx context.Context, jurusanKode int) jurusanWeb.JurusanWebResponse{
 tx,err := service.Db.Begin()
 helper.PanicIfError(err)
 defer helper.CommitOrRollback(tx)

 jurusan,err := service.JurusanRepository.FindJurusanByKode(ctx,tx,jurusanKode)
 if err != nil{
	panic(exception.NewNotFoundError(err.Error()))
 }

 return helper.ToJurusanResponse(jurusan)
}

func(service *JurusanServiceImpl)FindAllJurusan(ctx context.Context) []jurusanWeb.JurusanWebResponse	{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
 
	jurusan:= service.JurusanRepository.FindAllJurusan(ctx,tx)
	if err != nil{
	 helper.PanicIfError(err)
	}
 
	return helper.ToJurusanResponses(jurusan)
}
