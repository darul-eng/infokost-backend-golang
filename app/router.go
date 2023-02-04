package app

import (
	"backend-golang/controller"
	"backend-golang/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, boardingController controller.BoardingController, imageController controller.ImageController) *httprouter.Router {
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

	return router
}
