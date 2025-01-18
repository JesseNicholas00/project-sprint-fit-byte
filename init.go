package main

import (
	"github.com/JesseNicholas00/FitByte/controllers"
	activityCtrl "github.com/JesseNicholas00/FitByte/controllers/activity"
	imageCtrl "github.com/JesseNicholas00/FitByte/controllers/image"
	userCtrl "github.com/JesseNicholas00/FitByte/controllers/user"
	"github.com/JesseNicholas00/FitByte/middlewares"
	activityRepo "github.com/JesseNicholas00/FitByte/repos/activity"
	userRepo "github.com/JesseNicholas00/FitByte/repos/user"
	activitySvc "github.com/JesseNicholas00/FitByte/services/activity"
	userScv "github.com/JesseNicholas00/FitByte/services/user"
	"github.com/JesseNicholas00/FitByte/utils/ctxrizz"
	"github.com/JesseNicholas00/FitByte/utils/logging"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/jmoiron/sqlx"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
	uploader *manager.Uploader,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Error("panic while initializing controllers: %s", r)
			panic(r)
		}
	}()

	dbRizzer := ctxrizz.NewDbContextRizzer(db)

	userRepository := userRepo.NewUserRepository(dbRizzer)
	userService := userScv.NewUserService(userRepository, dbRizzer, cfg.jwtSecretKey, cfg.bcryptSaltCost, cfg.experimental)
	authMw := middlewares.NewAuthMiddleware(userService, cfg.experimental)
	userController := userCtrl.NewUserController(userService, authMw)
	ctrls = append(ctrls, userController)

	activityRepository := activityRepo.NewActivityRepository(dbRizzer)
	activityService := activitySvc.NewActivityService(activityRepository, dbRizzer)
	activityMw := middlewares.NewAuthMiddleware(userService, cfg.experimental)
	activityController := activityCtrl.NewActivityController(activityService, activityMw)
	ctrls = append(ctrls, activityController)

	imageController := imageCtrl.NewImageController(
		uploader,
		cfg.awsS3BucketName,
		cfg.awsRegion,
		authMw,
	)
	ctrls = append(ctrls, imageController)

	return
}
