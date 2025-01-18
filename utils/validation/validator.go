package validation

import (
	"reflect"
	"strings"

	types "github.com/JesseNicholas00/FitByte/types/optional"
	"github.com/JesseNicholas00/FitByte/utils/validation/iso8601"
	"github.com/JesseNicholas00/FitByte/utils/validation/optional"
	"github.com/JesseNicholas00/FitByte/utils/validation/uri"

	"github.com/JesseNicholas00/FitByte/utils/validation/image"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

var customFields = []customField{
	{
		Tag:       "imageExt",
		Validator: image.ValidateImageExtension,
	},
	{
		Tag:       "complete_uri",
		Validator: uri.ValidateCompleteURI,
	},
	{
		Tag:       "iso8601",
		Validator: iso8601.ValidateIso8601,
	},
}

type customField struct {
	Tag       string
	Validator validator.Func
}

var customTypes = []customType{
	{
		Type:      types.OptionalStr{},
		Validator: optional.ValidateOptionalString,
	},
	{
		Type:      types.OptionalUUID{},
		Validator: optional.ValidateOptionalUUID,
	},
	{
		Type:      types.OptionalInt{},
		Validator: optional.ValidateOptionalInt,
	},
}

type customType struct {
	Type      any
	Validator validator.CustomTypeFunc
}

func NewEchoValidator() echo.Validator {
	validator := validator.New(validator.WithRequiredStructEnabled())

	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	for _, customField := range customFields {
		validator.RegisterValidation(customField.Tag, customField.Validator)
	}

	for _, customType := range customTypes {
		validator.RegisterCustomTypeFunc(customType.Validator, customType.Type)
	}

	return &EchoValidator{
		validator: validator,
	}
}
