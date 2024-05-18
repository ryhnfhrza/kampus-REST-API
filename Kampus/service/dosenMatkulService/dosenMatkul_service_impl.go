package Service

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	web "kampus/model/web/dosenKelasMatkulWeb"
	"kampus/repository/dosenRepository"
	repository "kampus/repository/dosen_kelas_matkulRepository"
	"kampus/repository/matakuliahRepository"

	"github.com/go-playground/validator/v10"
)

type DosenKelasMatkulServiceImpl struct {
	DosenMatkulRepository repository.DosenKelasMatkulRepository
	Db *sql.DB
	Validate *validator.Validate
	DosenRepository dosenRepository.DosenRepository
	MatakuliahRepository matakuliahRepository.MatakuliahRepository
}

func NewDosenKelasMatkulService(dosenMatkulRepository repository.DosenKelasMatkulRepository,db *sql.DB,validate *validator.Validate,dosenRepository dosenRepository.DosenRepository,matakuliahRepository matakuliahRepository.MatakuliahRepository)DosenKelasMatkulService{
	return &DosenKelasMatkulServiceImpl{
		DosenMatkulRepository: dosenMatkulRepository,
		Db: db,
		Validate: validate,
		DosenRepository: dosenRepository,
		MatakuliahRepository: matakuliahRepository,

	}
}

func(Service *DosenKelasMatkulServiceImpl)AjarMatkul(ctx context.Context, request web.DosenKelasMatkulCreateRequest) web.DosenKelasMatkulResponse{
	err := Service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dosenRepository,err := Service.DosenRepository.FindById(ctx,tx,request.IdDosen)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	var dosenMatkuls []domain.DosenKelasMatkul
  var matakuliahs []domain.Matakuliah
    for _, kodeMatkul := range request.KodeMatkul {
        matakuliahRepository, err := Service.MatakuliahRepository.FindByKode(ctx, tx, kodeMatkul)
        if err != nil {
            panic(exception.NewNotFoundError(err.Error()))
        }
				matakuliahs = append(matakuliahs, matakuliahRepository)

        dosenMatkul := domain.DosenKelasMatkul{
            IdDosen: dosenRepository.Id,
            KodeMatkul: matakuliahRepository.Kode,
						KodeKelas: request.KodeKelas,
        }
        dosenMatkuls = append(dosenMatkuls, dosenMatkul)
       	Service.DosenMatkulRepository.AjarMatkul(ctx, tx, dosenMatkul)	
    }
		
    return helper.ToDosenMatkulResponse(dosenMatkuls, matakuliahs, dosenRepository,dosenMatkuls[0])
}