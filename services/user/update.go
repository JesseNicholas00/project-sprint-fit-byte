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
		user, err := svc.repo.FindUserByUserID(ctx, userID)
		switch {
		case err != nil:
			return errorutil.AddCurrentContext(err)
		}

		if req.Preference != nil {
			user.Preference.Scan(*req.Preference)
		}

		if req.WeightUnit != nil {
			user.WeightUnit.Scan(*req.WeightUnit)
		}

		if req.HeightUnit != nil {
			user.HeightUnit.Scan(*req.HeightUnit)
		}

		if req.Weight != nil {
			user.Weight = *req.Weight
		}

		if req.Height != nil {
			user.Height = *req.Height
		}

		if req.Name.V != nil {
			user.Name.Scan(*req.Name.V)
		}

		if req.ImageURI.V != nil {
			user.ImageURI.Scan(*req.ImageURI.V)
		}

		savedUser, err := svc.repo.UpdateUser(ctx, user)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = UpdateUserRes{
			Preference: savedUser.Preference.V,
			WeightUnit: savedUser.WeightUnit.V,
			HeightUnit: savedUser.HeightUnit.V,
			Weight:     savedUser.Weight,
			Height:     savedUser.Height,
			Email:      savedUser.Email,
			Name:       savedUser.Name.V,
			ImageURI:   savedUser.ImageURI.V,
		}

		return nil
	})
}
