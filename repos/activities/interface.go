package activity

import (
	"context"

	"github.com/google/uuid"
)

type ActivityRepository interface {
	AddActivity(ctx context.Context, activity Activity, userId uuid.UUID) error
}
