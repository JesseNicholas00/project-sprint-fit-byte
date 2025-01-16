package user

import (
	"errors"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *userController) loginUser(c echo.Context) error {
	req := user.AuthenticationUserReq{}
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	res := user.AuthenticationUserRes{}

	err := ctrl.service.LoginUser(c.Request().Context(), req, &res)
	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "kamu siapa?",
		})
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusOK, res)
}
