package activity

import (
	"net/http"

	"github.com/JesseNicholas00/FitByte/services/activity"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *activityController) addActivity(ctx echo.Context) error {
	req := activity.AddActivityReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		// log.Printf("error binding and validating: %v", err)
		return err
	}

	userId := ctx.Get("session").(user.GetSessionFromTokenRes).UserID

	res := activity.AddActivityRes{}

	if err := ctrl.service.AddActivity(ctx.Request().Context(), req, &res, userId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusCreated, res)
}
