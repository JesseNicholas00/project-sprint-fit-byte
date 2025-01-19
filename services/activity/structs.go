package activity

import (
	"errors"
	"github.com/JesseNicholas00/FitByte/types/optional"
	"github.com/google/uuid"
	"slices"
	"time"
)

type AddActivityReq struct {
	ActivityType      string    `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            string    `json:"doneAt" validate:"required,iso8601"`
	DurationInMinutes int       `json:"durationInMinutes" validate:"required,min=1"`
	UserID            uuid.UUID `json:"userId"`
}

func (_ AddActivityReq) BindBody() {}

func (r AddActivityReq) Validation() error {
	var errs error

	if !slices.Contains(validActivityTypes, r.ActivityType) {
		errs = errors.Join(errs, errors.New("ActivityType ga valid cuy"))
	}

	if _, err := time.Parse(time.RFC3339, r.DoneAt); err != nil {
		errs = errors.Join(errs, errors.New("DoneAt formatnya salah cuy"))
	}

	if r.DurationInMinutes < 1 {
		errs = errors.Join(errs, errors.New("DurationInMinutes minimal 1"))
	}

	return errs
}

var validActivityTypes = []string{"Walking", "Yoga", "Stretching", "Cycling", "Swimming", "Dancing", "Hiking", "Running", "HIIT", "JumpRope"}

type GetActivityResp []AddActivityRes

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
	DurationInMinutes optional.OptionalInt `json:"durationInMinutes" validate:"omitnil,min=1"`
}

func (_ UpdateActivityReq) BindBody() {}

func (r UpdateActivityReq) Validation() error {
	var errs error

	if r.ActivityType.Defined {
		switch {
		case r.ActivityType.V == nil:
			errs = errors.Join(errs, errors.New("ActivityType ga valid cuy"))
		case !slices.Contains(validActivityTypes, *r.ActivityType.V):
			errs = errors.Join(errs, errors.New("ActivityType ga valid cuy"))
		}
	}

	if r.DoneAt.Defined {
		switch {
		case r.DoneAt.V == nil:
			errs = errors.Join(errs, errors.New("DoneAt ga valid cuy"))
		default:
			if _, err := time.Parse(time.RFC3339, *r.DoneAt.V); err != nil {
				errs = errors.Join(errs, errors.New("DoneAt formatnya salah cuy"))
			}
		}
	}

	if r.DurationInMinutes.Defined {
		switch {
		case r.DurationInMinutes.V == nil:
			errs = errors.Join(errs, errors.New("DurationInMinutes ga valid cuy"))
		case *r.DurationInMinutes.V < 1:
			errs = errors.Join(errs, errors.New("DurationInMinutes minimal 1"))
		}
	}

	return errs
}

type GetActivityReq struct {
	Limit             *int   `query:"limit"`
	Offset            *int   `query:"offset"`
	ActivityType      string `query:"activityType"`
	DoneAtFrom        string `query:"doneAtFrom"`
	DoneAtTo          string `query:"doneAtTo"`
	CaloriesBurnedMin int    `query:"caloriesBurnedMin"`
	CaloriesBurnedMax int    `query:"caloriesBurnedMax"`
}
