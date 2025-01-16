package user

import (
	"errors"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *userController) registerUser(c echo.Context) error {
	req := user.AuthenticationUserReq{}
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	res := user.AuthenticationUserRes{}

	err := ctrl.service.RegisterUser(c.Request().Context(), req, &res)
	switch {
	case errors.Is(err, user.ErrEmailAlreadyRegistered):
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "login aja kan udah daftar tadi",
		})
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusCreated, res)
}
