package user

import (
	"errors"
	"github.com/JesseNicholas00/FitByte/types/optional"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/url"
	"slices"
	"strings"
)

type AuthenticationUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (_ AuthenticationUserReq) BindBody() {}

func (r AuthenticationUserReq) Validation() error {
	var errs error

	if !govalidator.IsEmail(r.Email) {
		errs = errors.Join(errs, errors.New("Ini sih bukan email"))
	}

	if len(r.Password) < 8 || len(r.Password) > 32 {
		errs = errors.Join(errs, errors.New("Passwordnya 8-32 karakter bwang"))
	}

	return errs
}

type UpdateUserReq struct {
	Name       optional.OptionalStr `json:"name" validate:"omitnil,min=2,max=60"`
	Preference string               `json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit string               `json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit string               `json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     int                  `json:"weight" validate:"required,min=10,max=1000"`
	Height     int                  `json:"height" validate:"required,min=3,max=250"`
	ImageURI   optional.OptionalStr `json:"imageUri" validate:"omitnil,complete_uri"`
}

func (_ UpdateUserReq) BindBody() {}

func (r UpdateUserReq) Validation() error {
	var errs error

	if r.Name.Defined {
		switch {
		case r.Name.V == nil:
			errs = errors.Join(errs, errors.New("Salah type om"))
		case len(*r.Name.V) < 2 || len(*r.Name.V) > 60:
			errs = errors.Join(errs, errors.New("Name itu harus 2-60 karakter"))
		}
	}

	if !slices.Contains(validPreferences, r.Preference) {
		errs = errors.Join(errs, errors.New("Preference nya ga valid"))
	}

	if !slices.Contains(validWeightUnits, r.WeightUnit) {
		errs = errors.Join(errs, errors.New("WeightUnit nya ga valid"))
	}

	if !slices.Contains(validHeightUnits, r.HeightUnit) {
		errs = errors.Join(errs, errors.New("HeightUnit nya ga valid"))
	}

	if r.Weight < 10 || r.Weight > 1000 {
		errs = errors.Join(errs, errors.New("Weight harus 10-1000"))
	}

	if r.Height < 3 || r.Height > 250 {
		errs = errors.Join(errs, errors.New("Height harus 3-250"))
	}

	if r.ImageURI.Defined {
		switch {
		case r.ImageURI.V == nil:
			errs = errors.Join(errs, errors.New("Salah type om"))
		default:
			imgURI, err := url.ParseRequestURI(*r.ImageURI.V)
			if err != nil {
				errs = errors.Join(errs, err)
			} else if !strings.Contains(imgURI.Host, ".") {
				errs = errors.Join(errs, errors.New("Salah url kayaknya om"))
			}
		}
	}

	return errs
}

var validPreferences = []string{"CARDIO", "WEIGHT"}

var validWeightUnits = []string{"KG", "LBS"}

var validHeightUnits = []string{"CM", "INCH"}

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
	Name       string `json:"name"`
	Email      string `json:"email"`
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     int    `json:"weight"`
	Height     int    `json:"height"`
	ImageURI   string `json:"imageUri"`
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
