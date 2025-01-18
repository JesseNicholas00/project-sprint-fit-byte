package activity

import (
	"context"

	"github.com/google/uuid"
)

type ActivityService interface {
	AddActivity(ctx context.Context, req AddActivityReq, res *AddActivityRes, userId uuid.UUID) error
}
