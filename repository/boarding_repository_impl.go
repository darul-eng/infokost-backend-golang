package repository

import (
	"backend-golang/helper"
	"backend-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BoardingRepositoryImpl struct {
}

func (repository *BoardingRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, boarding domain.Boarding) domain.Boarding {
	SQL := `INSERT INTO boarding("name", "description", "address", "contact", "price", "longlat") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, boarding.Name, boarding.Description, boarding.Address, boarding.Contact, boarding.Price, boarding.LongLat)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	boarding.Id = lastInsertId

	return boarding
}

func (repository *BoardingRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, boarding domain.Boarding) domain.Boarding {
	SQL := `UPDATE "boarding" SET "name"=$1, "description"=$2, "address"=$3, "contact"=$4, "price"=$5, "longlat"=$6 WHERE "id"=$7`
	_, err := tx.ExecContext(ctx, SQL, boarding.Name, boarding.Description, boarding.Address, boarding.Contact, boarding.Price, boarding.LongLat, boarding.Id)
	helper.PanicIfError(err)

	return boarding
}

func (repository *BoardingRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, boarding domain.Boarding) {
	SQL := `DELETE FROM "boarding" WHERE "id" = $1`
	_, err := tx.ExecContext(ctx, SQL, boarding.Id)
	helper.PanicIfError(err)
}

func (repository *BoardingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, boardingId int) (domain.Boarding, error) {
	SQL := `SELECT "id", "name", "description", "address", "contact", "price", "longlat" FROM "boarding" WHERE "id"=$1`
	rows, err := tx.QueryContext(ctx, SQL, boardingId)
	helper.PanicIfError(err)
	defer rows.Close()

	boarding := domain.Boarding{}
	if rows.Next() {
		err := rows.Scan(&boarding.Id, &boarding.Name, &boarding.Description, &boarding.Address, &boarding.Contact, &boarding.Price, &boarding.LongLat)
		helper.PanicIfError(err)

		return boarding, nil
	} else {
		return boarding, errors.New("boarding is not found")
	}
}

func (repository *BoardingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Boarding {
	SQL := `SELECT "id", "name", "description", "address", "contact", "price", "longlat" FROM "boarding"`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var boardings []domain.Boarding
	for rows.Next() {
		boarding := domain.Boarding{}
		err := rows.Scan(&boarding.Id, &boarding.Name, &boarding.Description, &boarding.Address, &boarding.Contact, &boarding.Price, &boarding.LongLat)
		helper.PanicIfError(err)

		boardings = append(boardings, boarding)
	}

	return boardings
}
