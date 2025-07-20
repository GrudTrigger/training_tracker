package validInput

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Login    string `json:"login" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Role     string `json:"role" validate:"required,min=0,max=2"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}
