package user_dto

type LoginResponse struct {
	FullName string
	Username string
	Email    string
	Role     string
	Token    string
}
