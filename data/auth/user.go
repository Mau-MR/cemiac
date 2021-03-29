package auth

type User struct {
	Name string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Mail string `json:"mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}


