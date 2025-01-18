package activity

import (
	"context"
	"time"

	"github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/helper"
	"github.com/google/uuid"
)

func (s *activityServiceImpl) AddActivity(ctx context.Context, req AddActivityReq, res *AddActivityRes, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	activityId, err := uuid.NewV7()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	var caloriesBurned int

	switch req.ActivityType {
	case "Walking", "Yoga", "Stretching":
		caloriesBurned = req.DurationInMinutes * 4
	case "Cycling", "Swimming", "Dancing":
		caloriesBurned = req.DurationInMinutes * 8
	case "Hiking", "Running", "HIIT", "JumpRope":
		caloriesBurned = req.DurationInMinutes * 10
	}

	res.CaloriesBurned = caloriesBurned

	doneAt, err := helper.MustParse(req.DoneAt)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	activity := activity.Activity{
		ActivityId:        activityId,
		ActivityType:      req.ActivityType,
		DoneAt:            doneAt,
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    res.CaloriesBurned,
		CreateAt:          time.Now(),
		UpdateAt:          time.Now(),
		UserID:            userId,
	}

	if err := s.repo.AddActivity(ctx, activity, userId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddActivityRes{
		ActivityId:        activityId.String(),
		ActivityType:      req.ActivityType,
		DoneAt:            req.DoneAt,
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    res.CaloriesBurned,
		CreateAt:          time.Now().Format(time.RFC3339),
		UpdateAt:          time.Now().Format(time.RFC3339),
	}

	return nil
}
