package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BoardingController interface {
	CreateBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindBoardingById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
