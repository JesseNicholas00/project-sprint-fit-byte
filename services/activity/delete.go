package activity

import (
	"context"
)

func (svc *activityServiceImpl) DeleteActivity(ctx context.Context, activityId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	err := svc.repo.DeleteActivity(ctx, activityId)
	if err != nil {
		return err
	}

	return nil
}
