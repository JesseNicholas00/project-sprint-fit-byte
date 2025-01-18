package activity

import "github.com/JesseNicholas00/FitByte/utils/ctxrizz"

type activityRepositoryImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewActivityRepository(dbRizzer ctxrizz.DbContextRizzer) ActivityRepository {
	return &activityRepositoryImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
