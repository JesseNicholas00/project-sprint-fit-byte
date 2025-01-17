package user

import (
	"context"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
)

func (repo *userRepositoryImpl) UpdateUser(
	ctx context.Context,
	user User,
) (res User, err error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return User{}, errorutil.AddCurrentContext(err)
	}

	rows, err := sess.NamedStmt(ctx, repo.statements.update).Queryx(user)
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

	return user, nil
}
