package mahasiswacontroller

import (
	"kampus/helper"
	"kampus/model/web"
	"kampus/model/web/mahasiswaWeb"
	"kampus/service/mahasiswaService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MahasiswaControllerImpl struct{
	mahasiswaService mahasiswaService.MahasiswaService
}

func(mahasiswaController *MahasiswaControllerImpl)Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaCreateRequest := mahasiswaWeb.MahasiswaCreateRequest{}
	helper.ReadFromRequestBody(request,&mahasiswaCreateRequest)

	mahasiswaResponse := mahasiswaController.mahasiswaService.Create(request.Context(),mahasiswaCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: mahasiswaResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(mahasiswaController *MahasiswaControllerImpl)Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaUpdateRequest := mahasiswaWeb.MahasiswaUpdateRequest{}
	helper.ReadFromRequestBody(request,&mahasiswaUpdateRequest)

	mahasiswaNIM := params.ByName("mahasiswaNIM")
	mahasiswaUpdateRequest.NIM = mahasiswaNIM

	mahasiswaResponse := mahasiswaController.mahasiswaService.Update(request.Context(),mahasiswaUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: mahasiswaResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(mahasiswaController *MahasiswaControllerImpl)Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaNIM := params.ByName("mahasiswaNIM")
	
	mahasiswaController.mahasiswaService.Delete(request.Context(),mahasiswaNIM)
	webResponse := web.WebResponse{
		Code: http.StatusNoContent,
		Status: "NO CONTENT",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(mahasiswaController *MahasiswaControllerImpl)FindByNIM(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaNIM := params.ByName("mahasiswaNIM")

	mahasiswaResponse := mahasiswaController.mahasiswaService.FindByNIM(request.Context(),mahasiswaNIM)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: mahasiswaResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(mahasiswaController *MahasiswaControllerImpl)FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	mahasiswaResponses := mahasiswaController.mahasiswaService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: mahasiswaResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
