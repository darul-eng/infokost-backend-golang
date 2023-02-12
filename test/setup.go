package test

import (
	"backend-golang/app"
	"backend-golang/controller"
	"backend-golang/middleware"
	"backend-golang/repository"
	"backend-golang/service"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func setupTestDB() *sql.DB {
	host := "127.0.0.1"
	port := "5432"
	user := "admin"
	password := "admin"
	dbname := "db_infokost_test"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	boardingRepository := repository.NewBoardingRepository()
	boardingService := service.NewBoardingService(boardingRepository, db, validate)
	boardingController := controller.NewBoardingController(boardingService)

	imageRepository := repository.NewImageRepository()
	imageService := service.NewImageService(imageRepository, boardingRepository, db, validate)
	imageController := controller.NewImageController(imageService)

	router := app.NewRouter(userController, boardingController, imageController)

	return middleware.NewAuthMiddleware(router)
}
