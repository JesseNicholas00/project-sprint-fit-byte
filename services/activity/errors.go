package activity

import "errors"

var (
	ErrActivityNotFound   = errors.New("activityService: activity not found")
	ErrActivityIdNotFound = errors.New("activityService: activity id not found")
)
