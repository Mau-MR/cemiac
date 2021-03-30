package jwt

import (
	"github.com/Mau-MR/cemiac/data/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTManager is json web token manager
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

//UserClaims is a custom JWT claims that contains some user information
type UserClaims struct {
	jwt.StandardClaims
	User *auth.User
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

