package activity

import (
	"net/http"

	"github.com/JesseNicholas00/FitByte/services/activity"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/labstack/echo/v4"
)

func (e *activityController) getActivityByFilters(ctx echo.Context) error {
	var params activity.GetActivityReq

	if err := request.BindAndValidate(ctx, &params); err != nil {
		return err
	}

	if params.Limit == nil || *params.Limit < 1 {
		params.Limit = new(int)
		*params.Limit = 5
	}

	if params.Offset == nil || *params.Offset < 1 {
		params.Offset = new(int)
		*params.Offset = 0
	}

	userId := ctx.Get("session").(user.GetSessionFromTokenRes).UserID

	res := make(activity.GetActivityResp, 0, *params.Limit)
	if err := e.service.GetActivityByFilters(ctx.Request().Context(), params, &res, userId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
