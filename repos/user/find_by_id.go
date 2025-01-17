package user

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (repo *userRepositoryImpl) FindUserByID(ctx context.Context, id string) (User, error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return User{}, errorutil.AddCurrentContext(err)
	}

	user := User{}

	err = sess.
		Stmt(ctx, repo.statements.findByID).
		QueryRowxContext(ctx, id).
		StructScan(&user)

	switch {
	case err != nil:
		return User{}, errorutil.AddCurrentContext(err)
	}

	return user, nil
}
