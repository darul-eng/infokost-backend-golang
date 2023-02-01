package boarding

type BoardingCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Address     string `validate:"required" json:"address"`
	Contact     string `validate:"required" json:"contact"`
	Price       int    `validate:"required" json:"price"`
	LongLat     string `validate:"required" json:"long_lat"`
}
