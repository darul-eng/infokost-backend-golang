package main

import (
	"backend-golang/app"
	"backend-golang/controller"
	"backend-golang/helper"
	"backend-golang/middleware"
	"backend-golang/repository"
	"backend-golang/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	db := app.NewDB()
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

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
