package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ImageController interface {
	CreateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteImageById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindImageById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
