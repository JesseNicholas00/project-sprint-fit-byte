package activity

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (repo *activityRepositoryImpl) DeleteActivity(ctx context.Context, activityId string) (err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.delete).
		ExecContext(
			ctx,
			activityId,
		)

	if a, _ := rows.RowsAffected(); a == 0 {
		err = errorutil.AddCurrentContext(ErrActivityIdNotFound)
		return err
	}

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}
	return nil
}
