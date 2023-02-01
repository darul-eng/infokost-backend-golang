package repository

import (
	"backend-golang/helper"
	"backend-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type ImageRepositoryImpl struct {
}

func NewImageRepository() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (repository *ImageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image {
	SQL := `INSERT INTO image("boarding_id", "name") VALUES ($1, $2) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, image.BoardingId, image.Name)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	image.Id = lastInsertId

	return image
}

func (repository *ImageRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image {
	SQL := `UPDATE image SET boarding_id=$1, name=$2 WHERE id=$3`
	_, err := tx.ExecContext(ctx, SQL, image.BoardingId, image.Name, image.Id)
	helper.PanicIfError(err)

	return image
}

func (repository *ImageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, image domain.Image) {
	SQL := `DELETE FROM image where id=$1`
	_, err := tx.ExecContext(ctx, SQL, image.Id)
	helper.PanicIfError(err)
}

func (repository *ImageRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx, boardingId int) {
	SQL := `DELETE FROM image where boarding_id=$1`
	_, err := tx.ExecContext(ctx, SQL, boardingId)
	helper.PanicIfError(err)
}

func (repository *ImageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, imageId int) (domain.Image, error) {
	SQL := `SELECT id, boarding_id, name FROM image WHERE id=$1`
	rows, err := tx.QueryContext(ctx, SQL, imageId)
	helper.PanicIfError(err)
	defer rows.Close()

	image := domain.Image{}
	if rows.Next() {
		err := rows.Scan(&image.Id, &image.BoardingId, &image.Name)
		helper.PanicIfError(err)

		return image, nil
	} else {
		return image, errors.New("image not found")
	}
}

func (repository *ImageRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, boardingId int) []domain.Image {
	SQL := `SELECT id, boarding_id, name FROM image WHERE boarding_id=$1`
	rows, err := tx.QueryContext(ctx, SQL, boardingId)
	helper.PanicIfError(err)
	defer rows.Close()

	var images []domain.Image
	for rows.Next() {
		image := domain.Image{}
		err := rows.Scan(&image.Id, &image.BoardingId, &image.Name)
		helper.PanicIfError(err)

		images = append(images, image)
	}

	return images
}
