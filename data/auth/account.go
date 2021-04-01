package auth

type Account struct{
	Mail string `json:"mail" validate:"required,email" firestore:"mail"`
	Password string `json:"password" validate:"required" firestore:"password"`
}
func NewAccount(mail,password string)*Account{
	return &Account{
		Mail: mail,
		Password: password,
	}
}


