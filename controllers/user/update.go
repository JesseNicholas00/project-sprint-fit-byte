package user

import (
	"net/http"

	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *userController) updateUser(c echo.Context) error {
	req := user.UpdateUserReq{}
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	userId := c.Get("session").(user.GetSessionFromTokenRes).UserID.String()

	res := user.UpdateUserRes{}

	err := ctrl.service.UpdateUser(c.Request().Context(), userId, req, &res)
	switch {
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusOK, res)
}
