package activity

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ActivityId        uuid.UUID `db:"activity_id"`
	ActivityType      string    `db:"activity_type"`
	DoneAt            time.Time `db:"done_at"`
	DurationInMinutes int       `db:"duration_in_minutes"`
	CaloriesBurned    int       `db:"calories_burned"`
	CreateAt          time.Time `db:"created_at"`
	UpdateAt          time.Time `db:"updated_at"`
	UserID            uuid.UUID `db:"user_id"`
}

type FilterActivity struct {
	Limit             int       `db:"limit"`
	Offset            int       `db:"offset"`
	ActivityType      string    `db:"activity_type"`
	DoneAtFrom        string    `db:"done_at_from"`
	DoneAtTo          string    `db:"done_at_to"`
	CaloriesBurnedMin int       `db:"calories_burned_min"`
	CaloriesBurnedMax int       `db:"calories_burned_max"`
	UserID            uuid.UUID `db:"user_id"`
}
