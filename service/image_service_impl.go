package service

import (
	"backend-golang/exception"
	"backend-golang/helper"
	"backend-golang/model/domain"
	"backend-golang/model/web/image"
	"backend-golang/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ImageServiceImpl struct {
	ImageRepository    repository.ImageRepository
	BoardingRepository repository.BoardingRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewImageService(imageRepository repository.ImageRepository, boardingRepository repository.BoardingRepository, DB *sql.DB, validate *validator.Validate) ImageService {
	return &ImageServiceImpl{ImageRepository: imageRepository, BoardingRepository: boardingRepository, DB: DB, Validate: validate}
}

func (service *ImageServiceImpl) Create(ctx context.Context, request image.ImageCreateRequest, boardingId int) image.ImageResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	image := domain.Image{
		BoardingId: boarding.Id,
		Name:       request.Name,
	}

	image = service.ImageRepository.Save(ctx, tx, image)

	return helper.ToImageResponse(image)
}

func (service *ImageServiceImpl) Update(ctx context.Context, request image.ImageUpdateRequest, boardingId int, imageId int) image.ImageResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	image.BoardingId = request.BoardingId
	image.Name = request.Name

	image = service.ImageRepository.Update(ctx, tx, image)

	return helper.ToImageResponse(image)

}

func (service *ImageServiceImpl) Delete(ctx context.Context, imageId int, boardingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ImageRepository.Delete(ctx, tx, image)
}

func (service *ImageServiceImpl) DeleteAll(ctx context.Context, boardingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ImageRepository.DeleteAll(ctx, tx, boarding.Id)
}

func (service *ImageServiceImpl) FindById(ctx context.Context, imageId int, boardingId int) image.ImageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToImageResponse(image)
}

func (service *ImageServiceImpl) FindAll(ctx context.Context, boardingId int) []image.ImageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	boarding, err := service.BoardingRepository.FindById(ctx, tx, boardingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	images := service.ImageRepository.FindAll(ctx, tx, boarding.Id)

	return helper.ToImageResponses(images)
}
