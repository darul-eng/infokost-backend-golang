package image

type ImageCreateRequest struct {
	BoardingId int    `json:"boarding_id"`
	Name       string `json:"name"`
}
