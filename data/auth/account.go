package auth

type Account struct{
	Mail string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}


