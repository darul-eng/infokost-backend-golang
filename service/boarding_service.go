package service

import (
	"backend-golang/model/web/boarding"
	"context"
)

type BoardingService interface {
	Create(ctx context.Context, request boarding.BoardingCreateRequest) boarding.BoardingResponse
	Update(ctx context.Context, request boarding.BoardingUpdateRequest) boarding.BoardingResponse
	Delete(ctx context.Context, boardingId int)
	FindById(ctx context.Context, boardingId int) boarding.BoardingResponse
	FindAll(ctx context.Context) []boarding.BoardingResponse
}
