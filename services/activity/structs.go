package activity

import (
	"github.com/JesseNicholas00/FitByte/types/optional"
	"github.com/google/uuid"
)

type AddActivityReq struct {
	ActivityType      string    `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            string    `json:"doneAt" validate:"required,iso8601"`
	DurationInMinutes int       `json:"durationInMinutes" validate:"required,min=1"`
	UserID            uuid.UUID `json:"userId"`
}

type AddActivityRes struct {
	ActivityId        string `json:"activityId"`
	ActivityType      string `json:"activityType"`
	DoneAt            string `json:"doneAt"`
	DurationInMinutes int    `json:"durationInMinutes"`
	CaloriesBurned    int    `json:"caloriesBurned"`
	CreateAt          string `json:"createdAt"`
	UpdateAt          string `json:"updatedAt"`
}

type UpdateActivityReq struct {
	ActivityType      optional.OptionalStr `json:"activityType" validate:"omitnil,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            optional.OptionalStr `json:"doneAt" validate:"omitnil,iso8601"`
	DurationInMinutes int                  `json:"durationInMinutes" validate:"omitnil,min=1"`
}
