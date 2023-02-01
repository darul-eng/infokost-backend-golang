package controller

import (
	"backend-golang/helper"
	"backend-golang/model/web"
	"backend-golang/model/web/image"
	"backend-golang/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ImageControllerImpl struct {
	ImageService service.ImageService
}

func NewImageController(imageService service.ImageService) ImageController {
	return &ImageControllerImpl{ImageService: imageService}
}

func (controller *ImageControllerImpl) CreateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageCreateRequest := image.ImageCreateRequest{}
	helper.ReadFromRequestBody(request, &imageCreateRequest)

	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	imageResponse := controller.ImageService.Create(request.Context(), imageCreateRequest, id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   imageResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ImageControllerImpl) UpdateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageUpdateRequest := image.ImageUpdateRequest{}
	helper.ReadFromRequestBody(request, &imageUpdateRequest)

	boardingId := params.ByName("boardingId")
	boarding, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	imageId := params.ByName("imageId")
	image, err := strconv.Atoi(imageId)
	helper.PanicIfError(err)

	imageResponse := controller.ImageService.Update(request.Context(), imageUpdateRequest, boarding, image)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   imageResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ImageControllerImpl) DeleteImageById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	boarding, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	imageId := params.ByName("imageId")
	image, err := strconv.Atoi(imageId)
	helper.PanicIfError(err)

	controller.ImageService.Delete(request.Context(), image, boarding)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ImageControllerImpl) DeleteImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	controller.ImageService.DeleteAll(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ImageControllerImpl) FindImageById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	boarding, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	imageId := params.ByName("imageId")
	image, err := strconv.Atoi(imageId)
	helper.PanicIfError(err)

	imageResponse := controller.ImageService.FindById(request.Context(), image, boarding)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   imageResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ImageControllerImpl) FindAllImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	imageResponses := controller.ImageService.FindAll(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   imageResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
