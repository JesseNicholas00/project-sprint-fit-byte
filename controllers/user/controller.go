package user

import (
	"github.com/JesseNicholas00/FitByte/controllers"
	"github.com/JesseNicholas00/FitByte/middlewares"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/labstack/echo/v4"
)

type userController struct {
	service user.UserService
	authMw  middlewares.Middleware
}

func NewUserController(service user.UserService, authMw middlewares.Middleware) controllers.Controller {
	return &userController{
		service: service,
		authMw:  authMw,
	}
}

func (ctrl *userController) Register(server *echo.Echo) error {
	server.POST("/v1/register", ctrl.registerUser)
	server.POST("/v1/login", ctrl.loginUser)

	userGroup := server.Group("/v1/user")
	userGroup.Use(ctrl.authMw.Process)
	userGroup.GET("", ctrl.findUser)
	userGroup.PATCH("", ctrl.updateUser)

	return nil
}
