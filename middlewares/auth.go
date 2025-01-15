package middlewares

import (
	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	binder *echo.DefaultBinder
}

func (mw *authMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func NewAuthMiddleware() Middleware {
	return &authMiddleware{
		binder: &echo.DefaultBinder{},
	}
}
