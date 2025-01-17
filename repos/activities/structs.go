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
