package activity

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (repo *activityRepositoryImpl) UpdateActivity(ctx context.Context, activity Activity, activityId, userID string) (res Activity, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.update).
		QueryxContext(ctx,
			activity.ActivityType,
			activity.DoneAt,
			activity.DurationInMinutes,
			activity.CaloriesBurned,
			activity.UpdateAt,
			activityId,
			userID,
		)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	return activity, nil
}
