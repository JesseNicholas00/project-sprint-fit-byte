package activity

import (
	"github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/ctxrizz"
)

type activityServiceImpl struct {
	repo     activity.ActivityRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewActivityService(
	repo activity.ActivityRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) ActivityService {
	return &activityServiceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
