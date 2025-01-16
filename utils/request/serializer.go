package request

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/labstack/echo/v4"
	"net/http"
)

// parser use bytedance/sonic is faster JSON parser library compare to standard library encoding/json.
// Package encoding/json is using a lot of reflection to encode/decode.
// Reflection is expensive, so use it less
var parser = sonic.ConfigFastest

// SonicSerializer implement echo.JSONSerializer interface for custom echo JSON serializer
type SonicSerializer struct{}

func (s SonicSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := parser.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (s SonicSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := parser.NewDecoder(c.Request().Body).Decode(i)

	// https://pkg.go.dev/github.com/bytedance/sonic#readme-print-error
	se := &decoder.SyntaxError{}
	me := &decoder.MismatchTypeError{}

	switch {
	case errors.As(err, &se):
		return echo.NewHTTPError(http.StatusBadRequest, se.Description())
	case errors.As(err, &me):
		return echo.NewHTTPError(http.StatusBadRequest, me.Description())
	}

	return err
}
