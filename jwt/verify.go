package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//Verify verifies the acces token string an return a use claim if it es valid
func (manager *JWTManager) Verify(accesToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accesToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				//to check that signing method is no modified
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(manager.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

