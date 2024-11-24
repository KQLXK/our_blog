package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type claims struct {
	UserId int64
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var SignMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

func CreateAccessToken(userid int64) (string, error) {
	Myclaims := claims{
		UserId: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(SignMethod, Myclaims)
	sign, err := token.SignedString(jwtSecret)
	return sign, err
}

func CreateRefreshToken(userid int64) (string, error) {
	Myclaims := claims{
		UserId: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(SignMethod, Myclaims)
	sign, err := token.SignedString(jwtSecret)
	return sign, err
}

func GetUserIdFromToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return -1, err
	}
	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims.UserId, nil
	}
	return -1, errors.New("invalid token")
}
