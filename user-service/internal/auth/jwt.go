package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("super-secret-key")

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {  //This provides the key used to verify signature

		//checking the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return secret, nil
    })
}