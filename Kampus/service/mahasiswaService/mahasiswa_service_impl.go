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
	MahasiswaRepository mahasiswaRepository.MahasiswaRepository
	Db *sql.DB
	Validate *validator.Validate
	
}

func NewMahasiswaService(mahasiswaRepository mahasiswaRepository.MahasiswaRepository,db *sql.DB,validate *validator.Validate)MahasiswaService{
	return &MahasiswaServiceImpl{
		MahasiswaRepository: mahasiswaRepository,
		Db: db,
		Validate: validate,
	}
}

func(Service *MahasiswaServiceImpl)Create(ctx context.Context, request mahasiswaWeb.MahasiswaCreateRequest)mahasiswaWeb.MahasiswaResponse{
	err := Service.Validate.Struct(request)
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
	lastInsertNim := Service.MahasiswaRepository.NIMTerakhirMahasiswaPadaJurusan(ctx,tx,request.KodeJurusan,tahunAngkatan)
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

	// Kelas logic
	var kodeKelas string
	var kelasGender string
	kelas := 1 
	//menentukan kelas berdasarkan gender
	if request.Gender == "pria" {
		kelasGender = "A"
		}else{
		kelasGender = "B"
	}
	

	for {
				kodeKelas = kelasGender + strconv.Itoa(kelas)
        adaKelas := Service.MahasiswaRepository.CheckKelas(ctx, tx, kodeKelas, request.KodeJurusan)
        if adaKelas {
            jmlMhsPadaKelas, err := Service.MahasiswaRepository.JumlahMhsPadaKelas(ctx, tx, kodeKelas, request.KodeJurusan)
            if err != nil {
                panic(exception.NewNotFoundError(err.Error()))
            }
            if jmlMhsPadaKelas < 30 {
                // Kelas ditemukan dan tidak penuh, keluar dari loop
                break
            }
        } else {
            // Jika kelas tidak ada, buat kelas baru
            Service.MahasiswaRepository.CreateKelas(ctx, tx, kodeKelas, request.KodeJurusan)
            // Setelah membuat kelas baru, keluar dari loop
            break
        }
        // Jika kelas ditemukan tetapi penuh, maka akan coba kelas selanjutnya
        kelas++
    }

	

	mahasiswa := domain.Mahasiswa{
		NIM: nim,
		Nama: request.Nama,
		Gender: request.Gender,
		TanggalLahir: parseTime,
		Semester: 1,
		KodeJurusan: request.KodeJurusan,
		Angkatan: tahunAngkatan,
		KodeKelas: kodeKelas,
	}

	mahasiswa = Service.MahasiswaRepository.Create(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)Update(ctx context.Context,request mahasiswaWeb.MahasiswaUpdateRequest) mahasiswaWeb.MahasiswaResponse{
	err := Service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.MahasiswaRepository.FindByNim(ctx,tx,request.NIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.KodeKelas == ""{
		request.KodeKelas = mahasiswa.KodeKelas
	}else{
		mahasiswa.KodeKelas = request.KodeKelas
	}

	if request.Semester == 0 {
		request.Semester = mahasiswa.Semester
	}else{
		mahasiswa.Semester = request.Semester
	}
	

	mahasiswa = Service.MahasiswaRepository.Update(ctx,tx,mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)Delete(ctx context.Context, mahasiswaNIM string){
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.MahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	Service.MahasiswaRepository.Delete(ctx,tx,mahasiswa)
}

func(Service *MahasiswaServiceImpl)FindByNIM(ctx context.Context,mahasiswaNIM string)mahasiswaWeb.MahasiswaResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa,err := Service.MahasiswaRepository.FindByNim(ctx,tx,mahasiswaNIM)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMahasiswaResponse(mahasiswa)
}

func(Service *MahasiswaServiceImpl)FindAll(ctx context.Context)[]mahasiswaWeb.MahasiswaResponse{
	tx,err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := Service.MahasiswaRepository.FindAll(ctx,tx)

	return helper.ToMahasiswaResponses(mahasiswa)
}

func (Service *MahasiswaServiceImpl) FindMahasiswaMatkulDosen(ctx context.Context, mahasiswaNIM string) mahasiswaWeb.MahasiswaMatkulDosenResponse {
	tx, err := Service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := Service.MahasiswaRepository.FindMatkuldanDosen(ctx, tx, mahasiswaNIM)
	if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMahasiswaMatkulDosenResponse(mahasiswa)
}


