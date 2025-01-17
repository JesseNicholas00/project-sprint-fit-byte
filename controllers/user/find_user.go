package user

import (
	"net/http"

	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/labstack/echo/v4"
)

func (ctrl *userController) findUser(c echo.Context) error {

	userId := c.Get("session").(user.GetSessionFromTokenRes).UserID.String()

	res := user.FindUserRes{}

	err := ctrl.service.FindUser(c.Request().Context(), userId, &res)
	switch {
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusCreated, res)
}
