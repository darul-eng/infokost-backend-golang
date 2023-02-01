package service

import (
	"backend-golang/exception"
	"backend-golang/helper"
	"backend-golang/model/domain"
	"backend-golang/model/web/boarding"
	"backend-golang/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type BoardingServiceImpl struct {
	BoardingRepository repository.BoardingRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBoardingService(boardingRepository repository.BoardingRepository, DB *sql.DB, validate *validator.Validate) BoardingService {
	return &BoardingServiceImpl{BoardingRepository: boardingRepository, DB: DB, Validate: validate}
}

func (service *BoardingServiceImpl) Create(ctx context.Context, request boarding.BoardingCreateRequest) boarding.BoardingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding := domain.Boarding{
		Name:        request.Name,
		Description: request.Description,
		Address:     request.Address,
		Contact:     request.Contact,
		Price:       request.Price,
		LongLat:     request.LongLat,
	}

	boarding = service.BoardingRepository.Save(ctx, tx, boarding)

	return helper.ToBoardingResponse(boarding)
}

func (service *BoardingServiceImpl) Update(ctx context.Context, request boarding.BoardingUpdateRequest) boarding.BoardingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	boarding.Name = request.Name
	boarding.Description = request.Description
	boarding.Address = request.Address
	boarding.Contact = request.Contact
	boarding.Price = request.Price
	boarding.LongLat = request.LongLat

	boarding = service.BoardingRepository.Update(ctx, tx, boarding)

	return helper.ToBoardingResponse(boarding)
}

func (service *BoardingServiceImpl) Delete(ctx context.Context, boardingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BoardingRepository.Delete(ctx, tx, boarding)
}

func (service *BoardingServiceImpl) FindById(ctx context.Context, boardingId int) boarding.BoardingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBoardingResponse(boarding)
}

func (service *BoardingServiceImpl) FindAll(ctx context.Context) []boarding.BoardingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boardings := service.BoardingRepository.FindAll(ctx, tx)

	return helper.ToBoardingResponses(boardings)
}
