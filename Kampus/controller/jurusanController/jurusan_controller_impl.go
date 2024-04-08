package jurusanController

import (
	"kampus/helper"
	"kampus/model/web"
	"kampus/model/web/jurusanWeb"
	"kampus/service/jurusanService"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type JurusanControllerImpl struct{
	JurusanService jurusanService.JurusanService
}

func NewJurusanController(jurusanService jurusanService.JurusanService)JurusanController{
	return &JurusanControllerImpl{
		JurusanService: jurusanService,
	}
}

func(controller *JurusanControllerImpl)AddJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
 jurusanCreateRequest := jurusanWeb.JurusanCreateRequest{}
 helper.ReadFromRequestBody(request,&jurusanCreateRequest)

 jurusanResponse := controller.JurusanService.AddJurusan(request.Context(),jurusanCreateRequest)
 webResponse := web.WebResponse{
	Code: http.StatusCreated,
	Status: "CREATED",
	Data: jurusanResponse,
 }

 helper.WriteToResponseBody(writer,webResponse)
}

func(controller *JurusanControllerImpl)UpdateNamaJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
jurusanUpdateRequest := jurusanWeb.JurusanUpdateRequest{}
helper.ReadFromRequestBody(request,&jurusanUpdateRequest)

jurusanKode := params.ByName("jurusanKode")
jurusanKodeInt,err := strconv.Atoi(jurusanKode)
helper.PanicIfError(err)
jurusanUpdateRequest.KodeJurusan = jurusanKodeInt

 jurusanResponse := controller.JurusanService.UpdateNamaJurusan(request.Context(),jurusanUpdateRequest)
 webResponse := web.WebResponse{
	Code: http.StatusCreated,
	Status: "CREATED",
	Data: jurusanResponse,
 }

 helper.WriteToResponseBody(writer,webResponse)
}

func(controller *JurusanControllerImpl)DeleteJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	jurusanKode := params.ByName("jurusanKode")
	kode,err := strconv.Atoi(jurusanKode)
	helper.PanicIfError(err)
	
	controller.JurusanService.DeleteJurusan(request.Context(),kode)
	webResponse := web.WebResponse{
		Code: http.StatusNoContent,
		Status: "NO CONTENT",
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *JurusanControllerImpl)FindJurusanByKode(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	jurusanKode := params.ByName("jurusanKode")
	kode,err := strconv.Atoi(jurusanKode)
	helper.PanicIfError(err)
	
	jurusanResponse := controller.JurusanService.FindJurusanByKode(request.Context(),kode)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: jurusanResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *JurusanControllerImpl)FindAllJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	jurusanResponse := controller.JurusanService.FindAllJurusan(request.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: jurusanResponse,
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}
