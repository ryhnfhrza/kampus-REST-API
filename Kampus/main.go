package main

import (
	"kampus/app"
	"kampus/controller/dosenController"
	controllerDosenMatkul "kampus/controller/dosenMatkulController"
	"kampus/controller/jurusanController"
	"kampus/controller/mahasiswaController"
	controllerMahasiswaMatkul "kampus/controller/mahasiswaMatkulController"
	"kampus/controller/matakuliahController"
	"kampus/exception"
	"kampus/helper"
	"kampus/repository/dosenRepository"
	repositoryDosenMatkul "kampus/repository/dosen_kelas_matkulRepository"
	"kampus/repository/jurusanRepository"
	"kampus/repository/mahasiswaRepository"
	repositoryMahasiswaMatkul "kampus/repository/mahasiswa_matakuliahRepository"
	"kampus/repository/matakuliahRepository"
	ServiceDosenMatkul "kampus/service/dosenMatkulService"
	"kampus/service/dosenService"
	"kampus/service/jurusanService"
	ServiceMahasiswaMatkul "kampus/service/mahasiswaMatkulService"
	"kampus/service/mahasiswaService"
	"kampus/service/matakuliahService"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDb()

	//Jurusan
	jurusanRepository := jurusanRepository.NewJurusanRepositoryImpl()
	jurusanService := jurusanService.NewJurusanService(jurusanRepository,db,validate)
	jurusanController := jurusanController.NewJurusanController(jurusanService)

	//Mahasiswa
	mahasiswaRepository := mahasiswaRepository.NewMahasiswaRepositoryImpl()
	mahasiswaService := mahasiswaService.NewMahasiswaService(mahasiswaRepository,db,validate)
	mahasiswaController := mahasiswaController.NewMahasiswaController(mahasiswaService)

	//dosen
	dosenRepository := dosenRepository.NewDosenRepositoryImpl()
	dosenService := dosenService.NewDosenService(dosenRepository,db,validate)
	dosenController := dosenController.NewDosenController(dosenService)

	//matakuliah
	matakuliahRepository := matakuliahRepository.NewMatakuliahRepositoryImpl()
	matakuliahService := matakuliahService.NewMatakuliahService(matakuliahRepository,db,validate)
	matakuliahController := matakuliahController.NewMatakuliahController(matakuliahService)

	//mahasiswaMatakuliah
	mahasiswaMatakuliahRepository := repositoryMahasiswaMatkul.NewMahasiswaMatkulRepositoryImpl()
	mahasiswaMatakuliahService := ServiceMahasiswaMatkul.NewMahasiswaMatkulService(mahasiswaMatakuliahRepository,db,validate,mahasiswaRepository,matakuliahRepository)
	mahasiswaMatakuliahController := controllerMahasiswaMatkul.NewMahasiswaMatkulController(mahasiswaMatakuliahService)

	//dosenKelasMatakuliah
	dosenKelasMatakuliahRepository := repositoryDosenMatkul.NewDosenKelasMatkulRepositoryImpl()
	dosenKelasMatakuliahService := ServiceDosenMatkul.NewDosenKelasMatkulService(dosenKelasMatakuliahRepository,db,validate,dosenRepository,matakuliahRepository)
	dosenKelasMatakuliahController := controllerDosenMatkul.NewDosenKelasMatkulController(dosenKelasMatakuliahService)
	router := httprouter.New()

	//Jurusan
	router.GET("/api/jurusan",jurusanController.FindAllJurusan)
	router.GET("/api/jurusan/:jurusanKode",jurusanController.FindJurusanByKode)
	router.POST("/api/jurusan",jurusanController.AddJurusan)
	router.PUT("/api/jurusan/:jurusanKode",jurusanController.UpdateNamaJurusan)
	router.DELETE("/api/jurusan/:jurusanKode",jurusanController.DeleteJurusan)

	//Mahasiswa 
	router.GET("/api/mahasiswa",mahasiswaController.FindAll)
	router.GET("/api/mahasiswa/:mahasiswaNIM",mahasiswaController.FindByNIM)
	router.GET("/api/mahasiswaMatkulDosen/:mahasiswaNIM",mahasiswaController.FindMatkulDosen)
	router.POST("/api/mahasiswa",mahasiswaController.Create)
	router.PUT("/api/mahasiswa/:mahasiswaNIM",mahasiswaController.Update)
	router.DELETE("/api/mahasiswa/:mahasiswaNIM",mahasiswaController.Delete)

	//dosen
	router.GET("/api/dosen",dosenController.FindAll)
	router.GET("/api/dosen/:dosenId",dosenController.FindById)
	router.POST("/api/dosen",dosenController.Create)
	router.PUT("/api/dosen/:dosenId",dosenController.Update)
	router.DELETE("/api/dosen/:dosenId",dosenController.Delete)

	//matakuliah
	router.GET("/api/matakuliah",matakuliahController.FindAll)
	router.GET("/api/matakuliah/:matakuliahKode",matakuliahController.FindByKode)
	router.POST("/api/matakuliah",matakuliahController.Create)
	router.PUT("/api/matakuliah/:matakuliahKode",matakuliahController.Update)
	router.DELETE("/api/matakuliah/:matakuliahKode",matakuliahController.Delete)

	//mahasiswaMatakuliah
	router.POST("/api/ambil-matakuliah/:mahasiswaNIM",mahasiswaMatakuliahController.AmbilMatkul)
	
	//dosenMatakuliah
	router.POST("/api/ajar-matakuliah/:dosenId",dosenKelasMatakuliahController.AjarMatkul)
	
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}