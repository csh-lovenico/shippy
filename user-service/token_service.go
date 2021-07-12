package main

import (
	pb "github.com/csh980717/shippy/user-service/proto/auth"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte("lovenico_252")

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (t *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (t *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := CustomClaims{
		user, jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "auth-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
