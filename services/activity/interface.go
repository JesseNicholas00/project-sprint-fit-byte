package activity

import (
	"context"

	"github.com/google/uuid"
)

type ActivityService interface {
	AddActivity(ctx context.Context, req AddActivityReq, res *AddActivityRes, userId uuid.UUID) error
	UpdateActivity(ctx context.Context, req UpdateActivityReq, res *AddActivityRes, activityId, userId string) error
	GetActivityByFilters(ctx context.Context, req GetActivityReq, res *GetActivityResp, userId uuid.UUID) error
	DeleteActivity(ctx context.Context, activityId string) error
}
