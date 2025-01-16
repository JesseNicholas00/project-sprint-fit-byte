package main

import (
	"github.com/JesseNicholas00/FitByte/controllers"
	userCtrl "github.com/JesseNicholas00/FitByte/controllers/user"
	"github.com/JesseNicholas00/FitByte/middlewares"
	userRepo "github.com/JesseNicholas00/FitByte/repos/user"
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
	authMw := middlewares.NewAuthMiddleware()

	userRepository := userRepo.NewUserRepository(dbRizzer)
	userService := userScv.NewUserService(userRepository, dbRizzer, cfg.jwtSecretKey, cfg.bcryptSaltCost, cfg.experimental)
	userController := userCtrl.NewUserController(userService, authMw)
	ctrls = append(ctrls, userController)

	return
}
