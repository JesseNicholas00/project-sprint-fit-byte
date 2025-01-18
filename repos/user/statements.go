package user

import (
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create       *sqlx.NamedStmt
	findByEmail  *sqlx.Stmt
	findByUserID *sqlx.Stmt
	update       *sqlx.NamedStmt
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
		findByUserID: statementutil.MustPrepare(`
			SELECT *
			FROM users
			WHERE user_id = $1
		`),
		update: statementutil.MustPrepareNamed(`
			UPDATE users
			SET
				preference = :preference,
				weight_unit = :weight_unit,
				height_unit = :height_unit,
				weight = :weight,
				height = :height,
				name = :name,
				image_uri = :image_uri
			WHERE user_id = :user_id
			RETURNING *
		`),
	}
}
