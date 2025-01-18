package activity

import (
	"context"

	"github.com/google/uuid"
)

type ActivityRepository interface {
	AddActivity(ctx context.Context, activity Activity, userId uuid.UUID) error
	FindActivityByActivityId(ctx context.Context, activityId, userId string) (Activity, error)
	UpdateActivity(ctx context.Context, activity Activity, activityId, userId string) (Activity, error)
}
