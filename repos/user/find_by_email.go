package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (repo *userRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (User, error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return User{}, errorutil.AddCurrentContext(err)
	}

	user := User{}

	err = sess.
		Stmt(ctx, repo.statements.findByEmail).
		QueryRowxContext(ctx, email).
		StructScan(&user)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return User{}, ErrEmailNotFound
	case err != nil:
		return User{}, errorutil.AddCurrentContext(err)
	}

	return user, nil
}
