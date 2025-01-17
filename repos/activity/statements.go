package activity

import (
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add             *sqlx.NamedStmt
	getByActivityId *sqlx.Stmt
	update          *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO activities (activity_id, user_id, activity_type, done_at, duration_in_minutes, calories_burned, created_at, updated_at)
			VALUES (:activity_id, :user_id, :activity_type, :done_at, :duration_in_minutes, :calories_burned, :created_at, :updated_at)
		`),
		getByActivityId: statementutil.MustPrepare(`
			SELECT * FROM activities WHERE activity_id = $1 AND user_id = $2
		`),
		update: statementutil.MustPrepare(`
			UPDATE activities
			SET activity_type = $1, done_at = $2, duration_in_minutes = $3, calories_burned = $4, updated_at = $5
			WHERE activity_id = $6 AND user_id = $7
			RETURNING *
		`),
	}
}
