package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTManager is json web token manager
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}
type UserInfo struct {
	Name    string `firestore:"name"`
	Role    string `firestore:"role"`
	Surname string `firestore:"surname"`
	ID      string `firestore:"id,omitempty"`
}

//UserClaims is a custom JWT claims that contains some user information
type UserClaims struct {
	jwt.StandardClaims
	UserInfo *UserInfo
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}
