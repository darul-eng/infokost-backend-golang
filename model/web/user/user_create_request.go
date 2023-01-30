package user

type UserCreateRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
