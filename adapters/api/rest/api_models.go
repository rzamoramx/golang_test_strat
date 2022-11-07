package rest

type RequestRegister struct {
	User     string `json:"user" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseRegister struct {
	Status  string `json:"status" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type ResponseLogin struct {
}
