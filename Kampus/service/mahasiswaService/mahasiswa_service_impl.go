package mahasiswaService

import (
	"context"
	"database/sql"
	"kampus/exception"
	"kampus/helper"
	"kampus/model/domain"
	"kampus/model/web/mahasiswaWeb"
	"kampus/repository/mahasiswaRepository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type MahasiswaServiceImpl struct{
	mahasiswaRepository mahasiswaRepository.MahasiswaRepository
	Db *sql.DB
	validate *validator.Validate
}

func NewMahasiswaService(MahasiswaRepository mahasiswaRepository.MahasiswaRepository,db *sql.DB,Validate *validator.Validate)MahasiswaService{
	return &MahasiswaServiceImpl{
		mahasiswaRepository: MahasiswaRepository,
		Db: db,
		validate: Validate,
	}
}

func(Service *MahasiswaServiceImpl)Create(ctx context.Context, request mahasiswaWeb.MahasiswaCreateRequest)mahasiswaWeb.MahasiswaResponse{
	err := Service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//set angkatan sesuai tahun ditambahkannya mahasiswa
	timeNow := time.Now()

	angkatan := time.Date(timeNow.Year(), 1, 1, 0, 0, 0, 0, timeNow.Location())
	tahunAngkatan := angkatan.Year()
	
	//ubah tipe time angkatan menjadi string	
	angkatanString := strconv.Itoa(angkatan.Year())
	
	//ubah tipe int jurusan menjadi string
	kodeJurusanStr := strconv.Itoa(request.KodeJurusan)


	//Create nim
	var strNim string;
	var jmlhMhs int
	lastInsertNim := Service.mahasiswaRepository.NIMTerakhirMahasiswaPadaJurusan(ctx,tx,request.KodeJurusan,tahunAngkatan)
	if lastInsertNim == ""{
		jmlhMhs = 0
	}else{
		nimTerakhir,err := strconv.Atoi(lastInsertNim)
		helper.PanicIfError(err) 
		templateNIM := strconv.Itoa(request.KodeJurusan)+ strconv.Itoa(tahunAngkatan)+"000"
		templateNIMInt,err := strconv.Atoi(templateNIM)
		helper.PanicIfError(err)
		jmlhMhs = nimTerakhir -  templateNIMInt
	}
	jmlhMhs = jmlhMhs + 1
	strJmlMhs := strconv.Itoa(jmlhMhs)

	if jmlhMhs < 10{
		strNim = "00" + strJmlMhs
	}else if jmlhMhs >= 10 && jmlhMhs < 100{
		strNim = "0" + strJmlMhs
	}else{
		strNim = strJmlMhs
	}

	nim := kodeJurusanStr+angkatanString+strNim

	//ubah inputan tanggal lahir menjadi time.Time
	formatter := "2006-01-02"
	birthDate := request.TanggalLahir
	parseTime , err := time.Parse(formatter,birthDate)
	helper.PanicIfError(err)


	mahasiswa := domain.Mahasiswa{
		NIM: nim,
		Nama: request.Nama,
		Gender: request.Gender,
		TanggalLahir: parseTime,
		Semester: 1,
		KodeJurusan: request.KodeJurusan,
		Angkatan: tahunAngkatan,
	}

	mahasiswa = Service.mahasiswaRepository.Create(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)Update(ctx context.Context,request mahasiswaWeb.MahasiswaUpdateRequest) mahasiswaWeb.MahasiswaResponse{
	err := Service.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.mahasiswaRepository.FindByNim(ctx,tx,request.NIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	mahasiswa.Semester	= request.Semester

	mahasiswa = Service.mahasiswaRepository.Update(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)Delete(ctx context.Context, mahasiswaNIM string){
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.mahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	Service.mahasiswaRepository.Delete(ctx,tx,mahasiswa)
}

func(Service *MahasiswaServiceImpl)FindByNIM(ctx context.Context,mahasiswaNIM string)mahasiswaWeb.MahasiswaResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.mahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)FindAll(ctx context.Context)[]mahasiswaWeb.MahasiswaResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := Service.mahasiswaRepository.FindAll(ctx,tx)

	return helper.ToMahasiswaResponses(mahasiswa)
}
