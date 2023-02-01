package repository

import (
	"backend-golang/model/domain"
	"context"
	"database/sql"
)

type ImageRepository interface {
	Save(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image
	Update(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image
	Delete(ctx context.Context, tx *sql.Tx, image domain.Image)
	DeleteAll(ctx context.Context, tx *sql.Tx, boardingId int)
	FindById(ctx context.Context, tx *sql.Tx, imageId int) (domain.Image, error)
	FindAll(ctx context.Context, tx *sql.Tx, boardingId int) []domain.Image
}
