package rest

type RequestLogin struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestRegister struct {
	User     string `json:"user" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Response
	Token string `json:"token"`
}
