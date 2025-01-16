package user

import (
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create      *sqlx.NamedStmt
	findByEmail *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO users (user_id, email, password)
			VALUES (:user_id, :email, :password)
		`),
		findByEmail: statementutil.MustPrepare(`
			SELECT user_id, email
			FROM users
			WHERE email = $1 LIMIT 1
		`),
	}
}
