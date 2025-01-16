package user

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/lib/pq"
)

func (repo *userRepositoryImpl) CreateUser(ctx context.Context, user User) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	_, err = sess.NamedStmt(ctx, repo.statements.create).ExecContext(ctx, user)

	pqErr := &pq.Error{}
	switch {
	case errors.As(err, &pqErr) && pqErr.Code == ErrDuplicateUnique:
		return ErrEmailExists
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
