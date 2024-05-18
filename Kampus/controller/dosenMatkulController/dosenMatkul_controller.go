package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DosenMatkulController interface {
	AjarMatkul(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}