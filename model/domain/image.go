package domain

type Image struct {
	Id         int    `json:"id"`
	BoardingId int    `json:"boarding_id"`
	Name       string `json:"name"`
}
