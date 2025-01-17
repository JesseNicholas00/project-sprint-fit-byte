package activity

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (r *activityRepositoryImpl) FindActivityByActivityId(ctx context.Context, activityId, userId string) (res Activity, err error) {

	if err := ctx.Err(); err != nil {
		return Activity{}, err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Activity{}, err
	}

	rows, err := sess.
		Stmt(ctx, r.statements.getByActivityId).
		QueryxContext(ctx, activityId, userId)

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		found = true
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	if !found {
		err = ErrActivityIdNotFound
		return
	}

	return
}
