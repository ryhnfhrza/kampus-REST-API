package controller

import (
	"kampus/helper"
	"kampus/model/web"
	webMahasiswaMatkul "kampus/model/web/mahasiswaMatakuliahWeb"
	Service "kampus/service/mahasiswaMatkulService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MahasiswaMatkulControllerImpl struct {
	MahasiswaMatkulService Service.MahasiswaMatkulService
}

func NewMahasiswaMatkulController(mahasiswaMatkulService Service.MahasiswaMatkulService)MahasiswaMatkulController{
	return &MahasiswaMatkulControllerImpl{
		MahasiswaMatkulService: mahasiswaMatkulService,
	}
}

func(Controller *MahasiswaMatkulControllerImpl)AmbilMatkul(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaMatkulCreateRequest := webMahasiswaMatkul.MahasiswaMatkulCreateRequest{}
	helper.ReadFromRequestBody(request,&mahasiswaMatkulCreateRequest)

	mahasiswaNIM := params.ByName("mahasiswaNIM")
	mahasiswaMatkulCreateRequest.NIM = mahasiswaNIM

	mahasiswaMatkulResponse := Controller.MahasiswaMatkulService.AmbilMatkul(request.Context(),mahasiswaMatkulCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: mahasiswaMatkulResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}