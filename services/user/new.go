package user

import (
	"github.com/JesseNicholas00/FitByte/repos/user"
	"github.com/JesseNicholas00/FitByte/utils/ctxrizz"
)

type userServiceImpl struct {
	repo            user.UserRepository
	dbRizzer        ctxrizz.DbContextRizzer
	jwtSecret       []byte
	bcryptCost      int
	useExperimental bool
}

func NewUserService(
	repo user.UserRepository,
	dbRizzer ctxrizz.DbContextRizzer,
	jwtSecret string,
	bcryptCost int,
	useExperimental bool,
) UserService {
	return &userServiceImpl{
		repo:            repo,
		dbRizzer:        dbRizzer,
		jwtSecret:       []byte(jwtSecret),
		bcryptCost:      bcryptCost,
		useExperimental: useExperimental,
	}
}
