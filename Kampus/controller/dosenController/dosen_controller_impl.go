package dosenController

import (
	"kampus/helper"
	"kampus/model/web"
	"kampus/model/web/dosenWeb"
	"kampus/service/dosenService"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type DosenControllerImpl struct {
	dosenService dosenService.DosenService
}

func(dosenController *DosenControllerImpl)Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenCreateRequest := dosenWeb.DosenCreateRequest{}
	helper.ReadFromRequestBody(request,&dosenCreateRequest)

	dosenResponse := dosenController.dosenService.Create(request.Context(),dosenCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: dosenResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(dosenController *DosenControllerImpl)Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenUpdateRequest := dosenWeb.DosenUpdateRequest{}
	helper.ReadFromRequestBody(request,&dosenUpdateRequest)

	dosenId := params.ByName("dosenId")
	id,err := strconv.Atoi(dosenId)
	helper.PanicIfError(err)
	dosenUpdateRequest.Id= id

	dosenResponse := dosenController.dosenService.Update(request.Context(),dosenUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: dosenResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(dosenController *DosenControllerImpl)Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenId := params.ByName("dosenId")
	id,err := strconv.Atoi(dosenId)
	helper.PanicIfError(err)

	dosenController.dosenService.Delete(request.Context(),id)
	webResponse := web.WebResponse{
		Code: http.StatusNoContent,
		Status: "NO CONTENT",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(dosenController *DosenControllerImpl)FindById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenId := params.ByName("dosenId")
	id,err := strconv.Atoi(dosenId)
	helper.PanicIfError(err)

	dosenResponse := dosenController.dosenService.FindById(request.Context(),id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: dosenResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(dosenController *DosenControllerImpl)FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	dosenResponses := dosenController.dosenService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: dosenResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
