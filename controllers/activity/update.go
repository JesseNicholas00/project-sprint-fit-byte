package activity

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/FitByte/services/activity"
	"github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/JesseNicholas00/FitByte/utils/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ctrl *activityController) updateActivity(ctx echo.Context) error {
	// Get path variable (activityId)
	activityId := ctx.Param("activityId")
	if _, err := uuid.Parse(activityId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "invalid id",
		})
	}

	// Get user ID (this is manager ID)
	userId := ctx.Get("session").(user.GetSessionFromTokenRes).UserID.String()

	// Validate data
	req := activity.UpdateActivityReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	res := activity.AddActivityRes{}

	// Update found activity
	if err := ctrl.service.UpdateActivity(ctx.Request().Context(), req, &res, activityId, userId); err != nil {
		switch {
		case errors.Is(err, activity.ErrActivityNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "activity not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
