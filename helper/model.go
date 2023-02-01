package helper

import (
	"backend-golang/model/domain"
	"backend-golang/model/web/boarding"
	"backend-golang/model/web/image"
	"backend-golang/model/web/user"
)

func ToUserResponse(data domain.User) user.UserResponse {
	return user.UserResponse{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
	}
}

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToBoardingResponse(data domain.Boarding) boarding.BoardingResponse {
	return boarding.BoardingResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Address:     data.Address,
		Contact:     data.Contact,
		Price:       data.Price,
		LongLat:     data.LongLat,
	}
}

func ToBoardingResponses(boardings []domain.Boarding) []boarding.BoardingResponse {
	var boardingResponses []boarding.BoardingResponse
	for _, boarding := range boardings {
		boardingResponses = append(boardingResponses, ToBoardingResponse(boarding))
	}

	return boardingResponses
}

func ToImageResponse(data domain.Image) image.ImageResponse {
	return image.ImageResponse{
		Id:         data.Id,
		BoardingId: data.BoardingId,
		Name:       data.Name,
	}
}

func ToImageResponses(images []domain.Image) []image.ImageResponse {
	var imageResponses []image.ImageResponse
	for _, image := range images {
		imageResponses = append(imageResponses, ToImageResponse(image))
	}

	return imageResponses
}
