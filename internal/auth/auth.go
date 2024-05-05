package auth

import (
	"campsite/internal/config"

	"github.com/dgrijalva/jwt-go"
)

func GetJWTsecretKey(config *config.Config) []byte {
	return []byte(config.JWT.Key)
}

// JWT Claims
type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}
