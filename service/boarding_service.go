package service

import (
	"backend-golang/model/web"
	"context"
)

type BoardingService interface {
	Create(ctx context.Context, request web.BoardingCreateRequest) web.BoardingResponse
	Update(ctx context.Context, request web.BoardingUpdateRequest) web.BoardingResponse
	Delete(ctx context.Context, boardingId int)
	FindById(ctx context.Context, boardingId int) web.BoardingResponse
	FIndAll(ctx context.Context) []web.BoardingResponse
}
