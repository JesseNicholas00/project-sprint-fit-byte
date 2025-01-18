package activity

import (
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO activities (activity_id, user_id, activity_type, done_at, duration_in_minutes, calories_burned, created_at, updated_at)
			VALUES (:activity_id, :user_id, :activity_type, :done_at, :duration_in_minutes, :calories_burned, :created_at, :updated_at)
		`),
	}
}
