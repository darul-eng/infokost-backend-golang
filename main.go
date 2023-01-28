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

	router := httprouter.New()

	router.GET("/api/user", userController.FindAll)
	router.POST("/api/user", userController.Create)
	router.GET("/api/user/:userId", userController.FindById)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler
	
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
