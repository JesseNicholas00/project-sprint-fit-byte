package iso8601

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func validateIso8601Impl(s string) bool {
	_, ok := time.Parse(time.RFC3339, s)
	return ok == nil
}

func ValidateIso8601(fl validator.FieldLevel) bool {
	return validateIso8601Impl(fl.Field().String())
}
