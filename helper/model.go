package helper

import (
	"backend-golang/model/domain"
	"backend-golang/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToBoardingResponse(boarding domain.Boarding) web.BoardingResponse {
	return web.BoardingResponse{
		Id:          boarding.Id,
		Name:        boarding.Name,
		Description: boarding.Description,
		Address:     boarding.Address,
		Contact:     boarding.Contact,
		Price:       boarding.Price,
		LongLat:     boarding.LongLat,
	}
}

func ToBoardingResponses(boardings []domain.Boarding) []web.BoardingResponse {
	var boardingResponses []web.BoardingResponse
	for _, boarding := range boardings {
		boardingResponses = append(boardingResponses, ToBoardingResponse(boarding))
	}

	return boardingResponses
}
