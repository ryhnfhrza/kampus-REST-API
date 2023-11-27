package mahasiswaController

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MahasiswaController interface {
	Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindByNIM(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}