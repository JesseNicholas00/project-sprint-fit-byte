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

type UpdateUserReq struct {
	Name       *string `json:"name" validate:"omitempty,min=2,max=60"`
	Preference *string `json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit *string `json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit *string `json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     *int    `json:"weight" validate:"required,min=10,max=1000"`
	Height     *int    `json:"height" validate:"required,min=3,max=250"`
	ImageURI   *string `json:"imageUri" validate:"omitempty,url"`
}

func (_ AuthenticationUserReq) BindBody() {}

func (r AuthenticationUserReq) Validation() error {
	var errs error

	if r.Email == "" {
		errs = errors.Join(errs, errors.New("Email ga bole kosong"))
	} else {
		if !govalidator.IsEmail(r.Email) {
			errs = errors.Join(errs, errors.New("Ini sih bukan email"))
		}
	}

	if r.Password == "" {
		errs = errors.Join(errs, errors.New("Password ga bole kosong"))
	} else {
		if len(r.Password) < 8 || len(r.Password) > 32 {
			errs = errors.Join(errs, errors.New("Passwordnya 8-32 karakter bwang"))
		}
	}

	return errs
}

type FindUserRes struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     int    `json:"weight"`
	Height     int    `json:"height"`
	ImageURI   string `json:"imageUri"`
}

type UpdateUserRes struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Preference *string `json:"preference"`
	WeightUnit *string `json:"weightUnit"`
	HeightUnit *string `json:"heightUnit"`
	Weight     *int    `json:"weight"`
	Height     *int    `json:"height"`
	ImageURI   *string `json:"imageUri"`
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
