package user

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/golang-jwt/jwt/v5"
)

func (svc *userServiceImpl) GetSessionFromToken(ctx context.Context, req GetSessionFromTokenReq, res *GetSessionFromTokenRes) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	claims := jwtClaims{}
	_, err := jwt.ParseWithClaims(
		req.Token,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return svc.jwtSecret, nil
		},
	)

	switch {
	case errors.Is(err, jwt.ErrTokenMalformed) ||
		errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return ErrTokenInvalid

	case errors.Is(err, jwt.ErrTokenExpired) ||
		errors.Is(err, jwt.ErrTokenNotValidYet):
		return ErrTokenExpired

	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	*res = GetSessionFromTokenRes{
		UserID: claims.Data.UserID,
	}

	return nil
}
