package user

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthenticationUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (_ AuthenticationUserReq) BindBody() {}

func (r AuthenticationUserReq) Validation() error {
	var errs error

	if r.Email == "" {
		errs = errors.Join(errs, errors.New("Email is required"))
	} else {
		if !govalidator.IsEmail(r.Email) {
			errs = errors.Join(errs, errors.New("Invalid email"))
		}
	}

	if r.Password == "" {
		errs = errors.Join(errs, errors.New("Password is required"))
	} else {
		if len(r.Password) < 8 || len(r.Password) > 32 {
			errs = errors.Join(errs, errors.New("Password must be between 8 and 32 characters"))
		}
	}

	return errs
}

type AuthenticationUserRes struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type GetSessionFromTokenReq struct {
	Token string
}

type GetSessionFromTokenRes struct {
	UserID uuid.UUID
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}

type jwtSubClaims struct {
	UserID uuid.UUID `json:"userId"`
}
