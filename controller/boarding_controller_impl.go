package controller

import (
	"backend-golang/helper"
	"backend-golang/model/web"
	"backend-golang/model/web/boarding"
	"backend-golang/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BoardingControllerImpl struct {
	BoardingService service.BoardingService
}

func NewBoardingController(boardingService service.BoardingService) BoardingController {
	return &BoardingControllerImpl{BoardingService: boardingService}
}

func (controller *BoardingControllerImpl) CreateBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingCreateRequest := boarding.BoardingCreateRequest{}
	helper.ReadFromRequestBody(request, &boardingCreateRequest)

	boardingResponse := controller.BoardingService.Create(request.Context(), boardingCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   boardingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BoardingControllerImpl) UpdateBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingUpdateRequest := boarding.BoardingUpdateRequest{}
	helper.ReadFromRequestBody(request, &boardingUpdateRequest)

	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	boardingUpdateRequest.Id = id

	boardingResponse := controller.BoardingService.Update(request.Context(), boardingUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   boardingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BoardingControllerImpl) DeleteBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	controller.BoardingService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BoardingControllerImpl) FindBoardingById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingId := params.ByName("boardingId")
	id, err := strconv.Atoi(boardingId)
	helper.PanicIfError(err)

	boardingResponse := controller.BoardingService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   boardingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BoardingControllerImpl) FindAllBoarding(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	boardingResponses := controller.BoardingService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   boardingResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
