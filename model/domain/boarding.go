package domain

type Boarding struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       []Image `json:"image"`
	Address     string  `json:"address"`
	Price       int     `json:"price"`
	LongLat     int     `json:"long_lat"`
}

type Image struct {
	Id         int    `json:"id"`
	BoardingId int    `json:"boarding_id"`
	Name       string `json:"name"`
}
