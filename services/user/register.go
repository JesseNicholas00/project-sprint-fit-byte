package user

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/FitByte/repos/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (svc *userServiceImpl) RegisterUser(ctx context.Context, req AuthenticationUserReq, res *AuthenticationUserRes) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	newUser := user.User{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
	}

	// EXPERIMENTAL:
	// Encryption is expensive, use if it is really necessary
	// Skip bcrypt encryption for cheaper compute
	if !svc.useExperimental {
		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			svc.bcryptCost,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		newUser.Password = string(cryptedPw)
	}

	err = svc.repo.CreateUser(ctx, newUser)
	switch {
	case errors.Is(err, user.ErrEmailExists):
		return ErrEmailAlreadyRegistered
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	token := id.String()

	// EXPERIMENTAL:
	// Encryption is expensive, use if it is really necessary
	// Skip jwt generation for cheaper compute
	if !svc.useExperimental {
		token, err = svc.generateToken(newUser)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
	}

	*res = AuthenticationUserRes{
		Email: req.Email,
		Token: token,
	}

	return nil
}
