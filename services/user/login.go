package user

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/FitByte/repos/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (svc *userServiceImpl) LoginUser(ctx context.Context, req AuthenticationUserReq, res *AuthenticationUserRes) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	loggedInUser, err := svc.repo.FindUserByEmail(ctx, req.Email)
	switch {
	case errors.Is(err, user.ErrEmailNotFound):
		return ErrUserNotFound
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	*res = AuthenticationUserRes{
		Email: loggedInUser.Email,
		Token: loggedInUser.ID.String(),
	}

	// EXPERIMENTAL:
	// Encryption is expensive, use if it is really necessary
	// Skip jwt generation for cheaper compute
	if !svc.useExperimental {
		res.Token, err = svc.generateToken(loggedInUser)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
	}

	return nil
}
