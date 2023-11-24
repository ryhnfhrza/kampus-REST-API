package mahasiswaService

import (
	"context"
	"database/sql"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/mahasiswaWeb"
	"kampus/repository/mahasiswaRepository"

	"github.com/go-playground/validator/v10"
)

type MahasiswaServiceImpl struct{
	mahasiswaRepository mahasiswaRepository.MahasiswaRepository
	Db *sql.DB
	validate *validator.Validate
}

func(mahasiswaService *MahasiswaServiceImpl)Create(ctx context.Context, request mahasiswaWeb.MahasiswaCreateRequest)mahasiswaWeb.MahasiswaResponse{
	err := mahasiswaService.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := mahasiswaService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := domain.Mahasiswa{
		NIM: request.NIM,
		Nama: request.Nama,
		Gender: request.Gender,
		Umur: request.Umur,
		Semester: request.Semester,
	}

	mahasiswa = mahasiswaService.mahasiswaRepository.Create(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(mahasiswaService *MahasiswaServiceImpl)Update(ctx context.Context,request mahasiswaWeb.MahasiswaUpdateRequest) mahasiswaWeb.MahasiswaResponse{
	err := mahasiswaService.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := mahasiswaService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := mahasiswaService.mahasiswaRepository.FindByNim(ctx,tx,request.NIM)
	helper.PanicIfError(err)

	mahasiswa.Nama = request.Nama
	mahasiswa.Semester	= request.Semester
	mahasiswa.Umur = request.Umur	

	mahasiswa = mahasiswaService.mahasiswaRepository.Update(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(mahasiswaService *MahasiswaServiceImpl)Delete(ctx context.Context, mahasiswaNIM string){
	tx,err := mahasiswaService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := mahasiswaService.mahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	helper.PanicIfError(err)

	mahasiswaService.mahasiswaRepository.Delete(ctx,tx,mahasiswa)
}

func(mahasiswaService *MahasiswaServiceImpl)FindByNIM(ctx context.Context,mahasiswaNIM string)mahasiswaWeb.MahasiswaResponse{
	tx,err := mahasiswaService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := mahasiswaService.mahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	helper.PanicIfError(err)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(mahasiswaService *MahasiswaServiceImpl)FindAll(ctx context.Context)[]mahasiswaWeb.MahasiswaResponse{
	tx,err := mahasiswaService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := mahasiswaService.mahasiswaRepository.FindAll(ctx,tx)

	return helper.ToMahasiswaResponses(mahasiswa)
}
