package repository

import (
	"backend-golang/model/domain"
	"context"
	"database/sql"
)

type BoardingRepository interface {
	Save(ctx context.Context, tx *sql.Tx, boarding domain.Boarding) domain.Boarding
	Update(ctx context.Context, tx *sql.Tx, boarding domain.Boarding) domain.Boarding
	Delete(ctx context.Context, tx *sql.Tx, boarding domain.Boarding)
	FindById(ctx context.Context, tx *sql.Tx, boardingId int) (domain.Boarding, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Boarding
}
