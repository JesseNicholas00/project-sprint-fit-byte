package activity

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/FitByte/repos/activity"
	"github.com/JesseNicholas00/FitByte/utils/errorutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ctrl *activityController) deleteActivity(ctx echo.Context) error {
	// Get path variable (activityId)
	activityId := ctx.Param("activityId")
	if _, err := uuid.Parse(activityId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "invalid id",
		})
	}

	// userId := ctx.Get("session").(user.GetSessionFromTokenRes).UserID.String()

	// Delete using service
	err := ctrl.service.DeleteActivity(ctx.Request().Context(), activityId)
	if err != nil {
		switch {
		case errors.Is(err, activity.ErrActivityIdNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "activity not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}
	return ctx.JSON(http.StatusOK, nil)
}
