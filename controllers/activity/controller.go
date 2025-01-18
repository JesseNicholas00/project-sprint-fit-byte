package activity

import (
	"github.com/JesseNicholas00/FitByte/controllers"
	"github.com/JesseNicholas00/FitByte/middlewares"
	"github.com/JesseNicholas00/FitByte/services/activity"
	"github.com/labstack/echo/v4"
)

type activityController struct {
	service activity.ActivityService
	authMw  middlewares.Middleware
}

func (ctrl *activityController) Register(server *echo.Echo) error {
	g := server.Group("/v1/activity", ctrl.authMw.Process) // Protected routes

	g.POST("", ctrl.addActivity)

	return nil
}

func NewActivityController(
	service activity.ActivityService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &activityController{
		service: service,
		authMw:  authMw,
	}
}
