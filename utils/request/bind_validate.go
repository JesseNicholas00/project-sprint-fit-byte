package request

import (
	"net/http"

	"github.com/JesseNicholas00/FitByte/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var echoDefaultBinder = &echo.DefaultBinder{}

func BindAndValidate[R any](
	ctx echo.Context,
	req *R,
) error {
	var err error

	if _, ok := any(req).(BodyBinder); ok {
		err = echoDefaultBinder.BindBody(ctx, req)
	} else {

		// echo.Context.Bind is parsing request with several steps.
		// Binding steps: 1) path params; 2) query params; 3) request body.
		// Step 1 and 2 using reflection which is expensive.
		// It is cheaper to bind data directly from specific source.
		// https://echo.labstack.com/docs/binding#direct-source
		err = ctx.Bind(req)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if reqValidator, ok := any(req).(Validator); ok {
		err = reqValidator.Validation()
	} else {

		// go-playground validator is using reflection which is expensive
		// It is cheaper but inconvenient to validate incoming request directly
		err = ctx.Validate(req)
	}
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": validation.ConvertToErrList(err),
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return nil
}

// BodyBinder signals BindAndValidate to used echo.DefaultBinder.BindBody directly
type BodyBinder interface {
	BindBody()
}

// Validator signals BindAndValidate to use request's Validation method to validate incoming request
type Validator interface {
	Validation() error
}
