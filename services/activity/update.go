package activity

import (
	"context"
	"time"

	repoActivity "github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/helper"
	"github.com/JesseNicholas00/FitByte/utils/transaction"
)

func (svc *activityServiceImpl) UpdateActivity(ctx context.Context, req UpdateActivityReq, res *AddActivityRes, activityId, userId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		activity, err := svc.repo.FindActivityByActivityId(ctx, activityId, userId)

		// Check if activity is not found
		if err != nil {
			switch err {
			case repoActivity.ErrActivityIdNotFound:
				return ErrActivityIdNotFound
			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		// Only process DoneAt if it's defined and has a value
		if req.DoneAt.Defined && req.DoneAt.V != nil {
			doneAt, err := helper.MustParse(*req.DoneAt.V)
			if err != nil {
				return errorutil.AddCurrentContext(err)
			}
			activity.DoneAt = doneAt
		}

		// Only update ActivityType if it's defined and has a value
		if req.ActivityType.Defined && req.ActivityType.V != nil {
			activity.ActivityType = *req.ActivityType.V
		}

		// Only update DurationInMinutes if it's defined and has a value
		if req.DurationInMinutes.Defined && req.DurationInMinutes.V != nil {
			activity.DurationInMinutes = int(*req.DurationInMinutes.V)
		}

		// Update calories burned based on activity type
		switch activity.ActivityType {
		case "Walking", "Yoga", "Stretching":
			activity.CaloriesBurned = activity.DurationInMinutes * 4
		case "Cycling", "Swimming", "Dancing":
			activity.CaloriesBurned = activity.DurationInMinutes * 8
		case "Hiking", "Running", "HIIT", "JumpRope":
			activity.CaloriesBurned = activity.DurationInMinutes * 10
		}

		// change updateAt to current time
		activity.UpdateAt, err = helper.MustParse(time.Now().Format(time.RFC3339))
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		result, err := svc.repo.UpdateActivity(ctx, activity, activityId, userId)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = AddActivityRes{
			ActivityId:        result.ActivityId.String(),
			ActivityType:      result.ActivityType,
			DoneAt:            result.DoneAt.Format(time.RFC3339),
			DurationInMinutes: result.DurationInMinutes,
			CaloriesBurned:    result.CaloriesBurned,
			CreateAt:          result.CreateAt.Format(time.RFC3339),
			UpdateAt:          result.UpdateAt.Format(time.RFC3339),
		}

		return nil
	})
}
