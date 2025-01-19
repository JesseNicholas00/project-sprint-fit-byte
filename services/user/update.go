package user

import (
	"context"
	"github.com/JesseNicholas00/FitByte/utils/helper"

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

		user.Preference = helper.Assign(req.Preference)
		user.WeightUnit = helper.Assign(req.WeightUnit)
		user.HeightUnit = helper.Assign(req.HeightUnit)
		user.Weight = req.Weight
		user.Height = req.Height

		if req.Name.Defined {
			user.Name = helper.Assign(*req.Name.V)
		}

		if req.ImageURI.V != nil {
			user.ImageURI = helper.Assign(*req.ImageURI.V)
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
