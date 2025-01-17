package helper

import "time"

func MustParse(timeString string) (time.Time, error) {
	res, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}

func MustParseDateOnly(timeString string) (time.Time, error) {
	res, err := time.Parse(time.DateOnly, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}
