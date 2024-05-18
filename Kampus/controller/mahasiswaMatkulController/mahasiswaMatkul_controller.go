package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MahasiswaMatkulController interface {
	AmbilMatkul(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}