package main

import (
	app "kampus/App"
	"kampus/controller/dosenController"
	"kampus/controller/mahasiswaController"
	"kampus/controller/matakuliahController"
	"kampus/exception"
	"kampus/helper"
	"kampus/repository/dosenRepository"
	"kampus/repository/mahasiswaRepository"
	"kampus/repository/matakuliahRepository"
	"kampus/service/dosenService"
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

	router := httprouter.New()

	//Mahasiswa 
	router.GET("/api/mahasiswa",mahasiswaController.FindAll)
	router.GET("/api/mahasiswa/:mahasiswaNIM",mahasiswaController.FindByNIM)
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

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}