package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthenticationUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationUserRes struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}

type jwtSubClaims struct {
	UserID uuid.UUID `json:"userId"`
}
