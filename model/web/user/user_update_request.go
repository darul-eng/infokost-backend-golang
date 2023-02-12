package user

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
