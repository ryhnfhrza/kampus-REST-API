package matakuliahController

import (
	"kampus/helper"
	"kampus/model/web"
	"kampus/model/web/matakuliahWeb"
	"kampus/service/matakuliahService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MatakuliahControllerImpl struct {
	matakuliahService matakuliahService.MatakuliahService
}

func NewMatakuliahController(Matakuliahservice matakuliahService.MatakuliahService)MatakuliahController{
	return &MatakuliahControllerImpl{
		matakuliahService: Matakuliahservice,
	}
}

func(matakuliahController *MatakuliahControllerImpl)Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	matakuliahCreateRequest := matakuliahWeb.MatakuliahCreateRequest{}
	helper.ReadFromRequestBody(request,&matakuliahCreateRequest)

	matakuliahResponse := matakuliahController.matakuliahService.Create(request.Context(),matakuliahCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: matakuliahResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(matakuliahController *MatakuliahControllerImpl)Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	matakuliahUpdateRequest := matakuliahWeb.MatakuliahUpdateRequest{}
	helper.ReadFromRequestBody(request,&matakuliahUpdateRequest)

	matakuliahKode := params.ByName("matakuliahKode")
	matakuliahUpdateRequest.Kode = matakuliahKode

	matakuliahResponse := matakuliahController.matakuliahService.Update(request.Context(),matakuliahUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: matakuliahResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(matakuliahController *MatakuliahControllerImpl)Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	matakuliahKode := params.ByName("matakuliahKode")

	matakuliahController.matakuliahService.Delete(request.Context(),matakuliahKode)
	webResponse := web.WebResponse{
		Code: http.StatusNoContent,
		Status: "NO CONTENT",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(matakuliahController *MatakuliahControllerImpl)FindByKode(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	matakuliahKode := params.ByName("matakuliahKode")

	matakuliahResponse := matakuliahController.matakuliahService.FindByKode(request.Context(),matakuliahKode)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: matakuliahResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(matakuliahController *MatakuliahControllerImpl)FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	matakuliahResponses := matakuliahController.matakuliahService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: matakuliahResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
