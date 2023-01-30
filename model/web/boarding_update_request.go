package web

type BoardingUpdateRequest struct {
	Id          int    `validate:"required" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	Price       int    `json:"price"`
	LongLat     int    `json:"long_lat"`
}
