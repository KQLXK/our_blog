package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"our_blog/dao"
	"time"
)

type claims struct {
	UserId int64
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var SignMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

func CreateUserToken(userid int64) (string, error) {
	accesstoken, err := CreateAccessToken(userid)
	if err != nil {
		return "", err
	}
	refreshtoken, err := CreateRefreshToken(userid)
	if err != nil {
		return "", err
	}
	err = dao.NewTokenDao().SetKey(accesstoken, refreshtoken)
	if err != nil {
		return "", err
	}
	return accesstoken, nil
}

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
	if err != nil {
		log.Println("create access token failed, err:", err)
	}
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
	if err != nil {
		log.Println("create refresh token failed, err:", err)
	}
	return sign, err
}

func GetUserIdFromToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Println("parse token failed, err:", err)
		return -1, err
	}
	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims.UserId, nil
	}
	return -1, errors.New("invalid token")
}

func ValidToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Println("parse token failed, err:", err)
		return false, err
	}
	if claims, ok := token.Claims.(jwt.RegisteredClaims); ok && token.Valid {
		if claims.ExpiresAt.Time.Before(time.Now()) {
			log.Println("token is expired")
			return true, nil
		}
		return false, nil
	}
	log.Println("invalid token")
	return true, nil
}
