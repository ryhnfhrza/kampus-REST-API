package matakuliahService

import (
	"context"
	"database/sql"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/matakuliahWeb"
	"kampus/repository/matakuliahRepository"

	"github.com/go-playground/validator/v10"
)

type matakuliahServiceImpl struct {
	matakuliahRepository matakuliahRepository.MatakuliahRepository
	Db *sql.DB
	validate *validator.Validate
}

func NewMatakuliahService(MatakuliahRepository matakuliahRepository.MatakuliahRepository,db *sql.DB,Validate *validator.Validate)MatakuliahService{
	return &matakuliahServiceImpl{
		matakuliahRepository: MatakuliahRepository,
		Db: db,
		validate: Validate,
	}
}

func(matakuliahService *matakuliahServiceImpl)Create(ctx context.Context, request matakuliahWeb.MatakuliahCreateRequest) matakuliahWeb.MatakuliahResponse{
	err := matakuliahService.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := matakuliahService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	matakuliah := domain.Matakuliah{
		Kode: request.Kode,
		Mata_kuliah: request.Matakuliah,
		SKS: request.SKS,
	}

	matakuliah = matakuliahService.matakuliahRepository.Create(ctx,tx,matakuliah)

	return helper.ToMatakuliahResponse(matakuliah)
}

func(matakuliahService *matakuliahServiceImpl)Update(ctx context.Context, request matakuliahWeb.MatakuliahUpdateRequest) matakuliahWeb.MatakuliahResponse{
	err := matakuliahService.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := matakuliahService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	matakuliah,err := matakuliahService.matakuliahRepository.FindByKode(ctx,tx,request.Kode)
	helper.PanicIfError(err)

	matakuliah.Mata_kuliah = request.Matakuliah
	matakuliah.SKS = request.SKS

	matakuliah = matakuliahService.matakuliahRepository.Update(ctx,tx,matakuliah)

	return helper.ToMatakuliahResponse(matakuliah)
}

func(matakuliahService *matakuliahServiceImpl)Delete(ctx context.Context, matakuliahKode string){
	tx,err := matakuliahService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	matakuliah,err := matakuliahService.matakuliahRepository.FindByKode(ctx,tx,matakuliahKode)
	helper.PanicIfError(err)

	matakuliahService.matakuliahRepository.Delete(ctx,tx,matakuliah)
}

func(matakuliahService *matakuliahServiceImpl)FindByKode(ctx context.Context, matakuliahKode string) matakuliahWeb.MatakuliahResponse{
	tx,err := matakuliahService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	matakuliah,err := matakuliahService.matakuliahRepository.FindByKode(ctx,tx,matakuliahKode)
	helper.PanicIfError(err)

	return helper.ToMatakuliahResponse(matakuliah)
}

func(matakuliahService *matakuliahServiceImpl)FindAll(ctx context.Context) []matakuliahWeb.MatakuliahResponse{
	tx,err := matakuliahService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	matakuliah := matakuliahService.matakuliahRepository.FindAll(ctx,tx)

	return helper.ToMatakuliahResponses(matakuliah)
}

