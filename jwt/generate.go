package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (manager *JWTManager) Generate(user *UserInfo) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		UserInfo: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}
