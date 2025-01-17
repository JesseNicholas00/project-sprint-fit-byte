package image

import (
	"github.com/JesseNicholas00/FitByte/controllers"
	"github.com/JesseNicholas00/FitByte/middlewares"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/labstack/echo/v4"
)

type imageController struct {
	service *manager.Uploader
	bucket  string
	region  string
	authMw  middlewares.Middleware
}

func (ctrl *imageController) Register(server *echo.Echo) error {
	g := server.Group("/v1/file")

	g.POST("", ctrl.uploadImage, ctrl.authMw.Process)

	return nil
}

func NewImageController(
	service *manager.Uploader,
	bucket string,
	region string,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &imageController{
		service: service,
		bucket:  bucket,
		region:  region,
		authMw:  authMw,
	}
}
