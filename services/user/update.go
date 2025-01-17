package user

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/transaction"
)

func (svc *userServiceImpl) UpdateUser(
	ctx context.Context,
	userID string,
	req UpdateUserReq,
	res *UpdateUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		user, err := svc.repo.FindUserByID(ctx, userID)
		switch {
		case err != nil:
			return errorutil.AddCurrentContext(err)
		}

		if req.Preference != nil {
			user.Preference.V = *req.Preference
		}

		if req.WeightUnit != nil {
			user.WeightUnit.V = *req.WeightUnit
		}

		if req.HeightUnit != nil {
			user.HeightUnit.V = *req.HeightUnit
		}

		if req.Weight != nil {
			user.Weight.V = *req.Weight
		}

		if req.Height != nil {
			user.Height.V = *req.Height
		}

		if req.Name != nil {
			user.Name.V = *req.Name
		}

		if req.ImageURI != nil {
			user.ImageURI.V = *req.ImageURI
		}

		_, err = svc.repo.UpdateUser(ctx, user)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = UpdateUserRes{
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
	})
}
