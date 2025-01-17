package activity

import (
	"context"
	"log"
	"time"

	"github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
	"github.com/shoenig/test/must"
)

func (s *activityServiceImpl) AddActivity(ctx context.Context, req AddActivityReq, res *AddActivityRes, userId uuid.UUID) error {
	log.Printf("service AddActivity: %v", req)
	if err := ctx.Err(); err != nil {
		return err
	}

	activityId := uuid.New()
	log.Printf("activityId: %v", activityId)

	switch req.ActivityType {
	case "Walking", "Yoga", "Stretching":
		res.CaloriesBurned = req.DurationInMinutes * 4
	case "Cycling", "Swimming", "Dancing":
		res.CaloriesBurned = req.DurationInMinutes * 8
	case "Hiking", "Running", "HIIT", "JumpRope":
		res.CaloriesBurned = req.DurationInMinutes * 10
	}

	activity := activity.Activity{
		ActivityId:        activityId,
		ActivityType:      req.ActivityType,
		DoneAt:            must.ParseTime(req.DoneAt),
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    res.CaloriesBurned,
		UserID:            userId,
	}

	if err := s.repo.AddActivity(ctx, activity, userId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddActivityRes{
		ActivityId:        activityId.String(),
		ActivityType:      req.ActivityType,
		DoneAt:            must.ParseTime(req.DoneAt),
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    res.CaloriesBurned,
		CreateAt:          time.Now().UTC().Format(time.RFC3339),
		UpdateAt:          time.Now().UTC().Format(time.RFC3339),
	}

	return nil
}
