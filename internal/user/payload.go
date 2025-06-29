package user

type CreateRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
