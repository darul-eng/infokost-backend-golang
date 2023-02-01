package main

import (
	"backend-golang/app"
	"backend-golang/controller"
	"backend-golang/exception"
	"backend-golang/helper"
	"backend-golang/middleware"
	"backend-golang/repository"
	"backend-golang/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()

	router.GET("/api/user", userController.FindAll)
	router.POST("/api/user", userController.Create)
	router.GET("/api/user/:userId", userController.FindById)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.GET("/api/boarding", boardingController.FindAllBoarding)
	router.POST("/api/boarding", boardingController.CreateBoarding)
	router.GET("/api/boarding/:boardingId", boardingController.FindBoardingById)
	router.PUT("/api/boarding/:boardingId", boardingController.UpdateBoarding)
	router.DELETE("/api/boarding/:boardingId", boardingController.DeleteBoarding)

	router.GET("/api/boarding/:boardingId/image", imageController.FindAllImage)
	router.POST("/api/boarding/:boardingId/image", imageController.CreateImage)
	router.DELETE("/api/boarding/:boardingId/image", imageController.DeleteImage)
	router.GET("/api/boarding/:boardingId/image/:imageId", imageController.FindImageById)
	router.PUT("/api/boarding/:boardingId/image/:imageId", imageController.UpdateImage)
	router.DELETE("/api/boarding/:boardingId/image/:imageId", imageController.DeleteImageById)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
