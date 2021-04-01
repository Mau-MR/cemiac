package jwt

type Token struct {
	Token string `json:"token"`
}

func NewToken(t string) *Token {
	return &Token{
		Token: t,
	}
}
