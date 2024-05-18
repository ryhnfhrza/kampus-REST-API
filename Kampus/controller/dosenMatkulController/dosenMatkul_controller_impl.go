package controller

import (
	"kampus/helper"
	"kampus/model/web"
	webDosenMatkul "kampus/model/web/dosenKelasMatkulWeb"
	Service "kampus/service/dosenMatkulService"
	"strconv"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DosenKelasMatkulControllerImpl struct {
	DosenKelasMatkulService Service.DosenKelasMatkulService
}

func NewDosenKelasMatkulController(dosenKelasMatkulService Service.DosenKelasMatkulService)DosenMatkulController{
	return &DosenKelasMatkulControllerImpl{
		DosenKelasMatkulService: dosenKelasMatkulService,
	}
}

func(Controller *DosenKelasMatkulControllerImpl)AjarMatkul(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenMatkulCreateRequest := webDosenMatkul.DosenKelasMatkulCreateRequest{}
	helper.ReadFromRequestBody(request,&dosenMatkulCreateRequest)

	dosenId := params.ByName("dosenId")
	dosenIdString,err := strconv.Atoi(dosenId)
	helper.PanicIfError(err)
	dosenMatkulCreateRequest.IdDosen = dosenIdString

	dosenMatkulResponse := Controller.DosenKelasMatkulService.AjarMatkul(request.Context(),dosenMatkulCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: dosenMatkulResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}