package user

import "github.com/JesseNicholas00/FitByte/utils/ctxrizz"

type userRepositoryImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewUserRepository(dbRizzer ctxrizz.DbContextRizzer) UserRepository {
	return &userRepositoryImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
