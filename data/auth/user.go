package auth

type User struct {
	Name     string `json:"name" validate:"required" firestore:"name,omitempty"`
	Surname  string `json:"surname" validate:"required" firestore:"surname,omitempty"`
	Mail     string `json:"mail" validate:"required" firestore:"mail,omitempty"`
	Password string `json:"password" validate:"required" firestore:"password,omitempty"`
	Role     string `json:"role"  validate:"required" firestore:"role,omitempty"`
}

func NewUser(name, surname, mail, password, role string) *User {
	return &User{
		Name:     name,
		Surname:  surname,
		Mail:     mail,
		Password: password,
		Role:     role,
	}
}
