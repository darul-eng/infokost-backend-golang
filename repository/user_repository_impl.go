package repository

import (
	"backend-golang/helper"
	"backend-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO "user"("name", "email", "password") VALUES ($1, $2, $3) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, user.Name, user.Email, user.Password)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	user.Id = lastInsertId

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `UPDATE "user" SET "name"=$1, "email"=$2, "password"=$3 WHERE "id"=$4`
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := `DELETE FROM "user" WHERE "id" = $1`
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := `SELECT "id", "name", "email" FROM "user" WHERE "id" = $1`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `SELECT "id", "name", "email" FROM "user"`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}
