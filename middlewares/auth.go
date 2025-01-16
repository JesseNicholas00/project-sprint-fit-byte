package middlewares

import (
	"errors"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type authMiddleware struct {
	service         user.UserService
	binder          *echo.DefaultBinder
	useExperimental bool
}

func (mw *authMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := struct {
			Bearer string `header:"Authorization"`
		}{}

		if err := mw.binder.BindHeaders(c, &header); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "invalid request",
			})
		}

		if header.Bearer == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "missing bearer token",
			})
		}

		splitByBearer := strings.Split(header.Bearer, "Bearer ")
		if len(splitByBearer) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "malformed bearer token",
			})
		}

		token := splitByBearer[1]
		req := user.GetSessionFromTokenReq{
			Token: token,
		}
		res := user.GetSessionFromTokenRes{}

		if mw.useExperimental {
			id, err := uuid.Parse(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"message": "malformed bearer token",
				})
			}

			res.UserID = id
		} else {
			if err := mw.service.GetSessionFromToken(
				c.Request().Context(),
				req,
				&res,
			); err != nil {
				switch {
				case errors.Is(err, user.ErrTokenInvalid):
					return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
						"message": "malformed bearer token",
					})

				case errors.Is(err, user.ErrTokenExpired):
					return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
						"message": "expired bearer token",
					})

				default:
					return errorutil.AddCurrentContext(err)
				}
			}
		}

		c.Set("session", res)

		return next(c)
	}
}

func NewAuthMiddleware(service user.UserService, useExperimental bool) Middleware {
	return &authMiddleware{
		service:         service,
		binder:          &echo.DefaultBinder{},
		useExperimental: useExperimental,
	}
}
