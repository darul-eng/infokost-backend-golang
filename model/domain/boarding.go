package domain

type Boarding struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       []Image `json:"image"`
	Address     string  `json:"address"`
	Contact     string  `json:"contact"`
	Price       int     `json:"price"`
	LongLat     string  `json:"long_lat"`
}
