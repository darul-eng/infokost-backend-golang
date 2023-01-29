package web

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
