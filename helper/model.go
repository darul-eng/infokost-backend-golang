package helper

import (
	"backend-golang/model/domain"
	"backend-golang/model/web/boarding"
	"backend-golang/model/web/user"
)

func ToUserResponse(user domain.User) user.UserResponse {
	return user.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToBoardingResponse(boarding domain.Boarding) boarding.BoardingResponse {
	return boarding.BoardingResponse{
		Id:          boarding.Id,
		Name:        boarding.Name,
		Description: boarding.Description,
		Address:     boarding.Address,
		Contact:     boarding.Contact,
		Price:       boarding.Price,
		LongLat:     boarding.LongLat,
	}
}

func ToBoardingResponses(boardings []domain.Boarding) []boarding.BoardingResponse {
	var boardingResponses []boarding.BoardingResponse
	for _, boarding := range boardings {
		boardingResponses = append(boardingResponses, ToBoardingResponse(boarding))
	}

	return boardingResponses
}
