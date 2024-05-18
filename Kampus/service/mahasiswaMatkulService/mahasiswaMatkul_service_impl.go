package Service

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	web "kampus/model/web/mahasiswaMatakuliahWeb"
	"kampus/repository/mahasiswaRepository"
	repository "kampus/repository/mahasiswa_matakuliahRepository"
	"kampus/repository/matakuliahRepository"

	"github.com/go-playground/validator/v10"
)

type MahasiswaMatkulServiceImpl struct {
	MahasiswaMatkulRepository repository.MahasiswaMatkulRepository
	Db *sql.DB
	validate *validator.Validate
	MahasiswaRepository mahasiswaRepository.MahasiswaRepository
	MatakuliahRepository matakuliahRepository.MatakuliahRepository
}

func NewMahasiswaMatkulService(mahasiswaMatkulRepository repository.MahasiswaMatkulRepository,db *sql.DB,Validate *validator.Validate,mahasiswaRepository mahasiswaRepository.MahasiswaRepository,matakuliahRepository matakuliahRepository.MatakuliahRepository)MahasiswaMatkulService{
	return &MahasiswaMatkulServiceImpl{
		MahasiswaMatkulRepository: mahasiswaMatkulRepository,
		Db: db,
		validate: Validate,
		MahasiswaRepository: mahasiswaRepository,
		MatakuliahRepository: matakuliahRepository,
	}
}

func(Service *MahasiswaMatkulServiceImpl)AmbilMatkul(ctx context.Context, request web.MahasiswaMatkulCreateRequest) web.MahasiswaMatkulResponse{
	err := Service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswaRepository,err := Service.MahasiswaRepository.FindByNim(ctx,tx,request.NIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	var mahasiswaMatkuls []domain.MahasiswaMatkul
  var matakuliahs []domain.Matakuliah
    for _, kodeMatkul := range request.KodeMatkul {
        matakuliahRepository, err := Service.MatakuliahRepository.FindByKode(ctx, tx, kodeMatkul)
        if err != nil {
            panic(exception.NewNotFoundError(err.Error()))
        }
				matakuliahs = append(matakuliahs, matakuliahRepository)

        mahasiswaMatkul := domain.MahasiswaMatkul{
            NIM:        mahasiswaRepository.NIM,
            KodeMatkul: matakuliahRepository.Kode,
        }
        mahasiswaMatkuls = append(mahasiswaMatkuls, mahasiswaMatkul)
       	Service.MahasiswaMatkulRepository.AmbilMatkul(ctx, tx, mahasiswaMatkul)	
    }
		
    return helper.ToMahasiswaMatkulResponse(mahasiswaMatkuls, matakuliahs, mahasiswaRepository)
}