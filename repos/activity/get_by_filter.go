package activity

import (
	"context"
	"github.com/JesseNicholas00/FitByte/utils/statementutil"
	"time"

	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/mewsql"
)

func (repo *activityRepositoryImpl) GetActivityByFilters(ctx context.Context, filter FilterActivity) ([]Activity, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var conditions []mewsql.Condition

	if filter.ActivityType != "" {
		conditions = append(conditions,
			mewsql.WithCondition("activity_type = ?", filter.ActivityType),
		)
	}

	if filter.DoneAtFrom != "" {
		doneAtFrom, err := time.Parse(time.RFC3339, filter.DoneAtFrom)
		if err != nil {
			return nil, errorutil.AddCurrentContext(err)
		}
		conditions = append(conditions,
			mewsql.WithCondition("done_at < ?", doneAtFrom),
		)
	}

	if filter.DoneAtTo != "" {
		doneAtTo, err := time.Parse(time.RFC3339, filter.DoneAtTo)
		if err != nil {
			return nil, errorutil.AddCurrentContext(err)
		}
		conditions = append(conditions,
			mewsql.WithCondition("done_at > ?", doneAtTo),
		)
	}

	if filter.CaloriesBurnedMin != 0 {
		conditions = append(conditions,
			mewsql.WithCondition("calories_burned >= ?", filter.CaloriesBurnedMin),
		)
	}

	if filter.CaloriesBurnedMax != 0 {
		conditions = append(conditions,
			mewsql.WithCondition("calories_burned <= ?", filter.CaloriesBurnedMax),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithLimit(filter.Limit),
		mewsql.WithOffset(filter.Offset),
		mewsql.WithWhere(conditions...),
	}

	sql, args := mewsql.Select(
		`activity_id, user_id, activity_type, done_at, calories_burned`,
		"activities",
		options...,
	)

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	stmt := statementutil.GetCachedStmt(sql)
	rows, err := sess.Stmt(ctx, stmt).QueryxContext(ctx, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	defer rows.Close()

	activities := make([]Activity, 0, filter.Limit)
	for rows.Next() {
		var activity Activity
		err = rows.StructScan(&activity)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}
