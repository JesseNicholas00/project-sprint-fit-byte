package user

import (
	"github.com/JesseNicholas00/FitByte/repos/user"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (svc *userServiceImpl) generateToken(user user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
		Data: jwtSubClaims{
			UserID: user.ID,
		},
	})

	return token.SignedString(svc.jwtSecret)
}
