package matakuliahController

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MatakuliahController interface {
	Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindByKode(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}