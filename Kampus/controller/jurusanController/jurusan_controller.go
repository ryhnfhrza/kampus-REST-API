package jurusanController

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JurusanController interface{
	AddJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	UpdateNamaJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	DeleteJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindJurusanByKode(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindAllJurusan(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}