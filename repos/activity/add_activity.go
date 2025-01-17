package activity

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
)

func (repo *activityRepositoryImpl) AddActivity(ctx context.Context, activity Activity, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	_, err = sess.NamedStmt(ctx, repo.statements.add).Exec(activity)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	return nil
}
