package jwt

import (
	"github.com/Mau-MR/cemiac/data/auth"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (manager *JWTManager) Generate(user *auth.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

