package auth

import (
	"test-case-majoo/entity/auths"
	"test-case-majoo/entity/responses"
)

//Writer user writer
type Writer interface {
}

type Reader interface {
	GetUserByID(userID string) (auths.User, *responses.Response)
	GetMonthlyReport(userID, Month string) (auths.MonthlyReport, *responses.Response)
}

//Repository interface
type Repository interface {
	Writer
	Reader
}

type UseCase interface {
	GetUserByID(userID string) (auths.User, *responses.Response)
	GetMonthlyReport(userID, Month string) (auths.MonthlyReport, *responses.Response)
}
