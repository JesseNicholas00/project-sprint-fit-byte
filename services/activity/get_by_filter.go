package activity

import (
	"context"
	"time"

	"github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
)

func (a *activityServiceImpl) GetActivityByFilters(ctx context.Context, params GetActivityReq, res *GetActivityResp, userId uuid.UUID) error {

	if err := ctx.Err(); err != nil {
		return err
	}

	activities, err := a.repo.GetActivityByFilters(ctx, activity.FilterActivity{
		Limit:             *params.Limit,
		Offset:            *params.Offset,
		ActivityType:      params.ActivityType,
		DoneAtFrom:        params.DoneAtFrom,
		DoneAtTo:          params.DoneAtTo,
		CaloriesBurnedMin: params.CaloriesBurnedMin,
		CaloriesBurnedMax: params.CaloriesBurnedMax,
		UserID:            userId,
	})

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, act := range activities {
		*res = append(*res, AddActivityRes{
			ActivityId:        act.ActivityId.String(),
			ActivityType:      act.ActivityType,
			DoneAt:            act.DoneAt.Format(time.RFC3339),
			DurationInMinutes: act.DurationInMinutes,
			CaloriesBurned:    act.CaloriesBurned,
			CreateAt:          act.CreateAt.Format(time.RFC3339),
			UpdateAt:          act.UpdateAt.Format(time.RFC3339),
		})
	}

	return nil
}
