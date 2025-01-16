package user

import (
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO users (user_id, email, password)
			VALUES (:user_id, :email, :password)
		`),
	}
}
