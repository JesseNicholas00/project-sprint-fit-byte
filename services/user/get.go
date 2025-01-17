package user

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (svc *userServiceImpl) FindUser(
	ctx context.Context,
	userID string,
	res *FindUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	user, err := svc.repo.FindUserByID(ctx, userID)
	switch {
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	*res = FindUserRes{
		Preference: &user.Preference.V,
		WeightUnit: &user.WeightUnit.V,
		HeightUnit: &user.HeightUnit.V,
		Weight:     &user.Weight.V,
		Height:     &user.Height.V,
		Email:      &user.Email,
		Name:       &user.Name.V,
		ImageURI:   &user.ImageURI.V,
	}

	return nil
}
