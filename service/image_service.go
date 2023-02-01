package service

import (
	"backend-golang/model/web/image"
	"context"
)

type ImageService interface {
	Create(ctx context.Context, request image.ImageCreateRequest, boardingId int) image.ImageResponse
	Update(ctx context.Context, request image.ImageUpdateRequest, boardingId int, imageId int) image.ImageResponse
	Delete(ctx context.Context, imageId int, boardingId int)
	DeleteAll(ctx context.Context, boardingId int)
	FindById(ctx context.Context, imageId int, boardingId int) image.ImageResponse
	FindAll(ctx context.Context, boardingId int) []image.ImageResponse
}
