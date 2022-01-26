package auth

import (
	"test-case-majoo/entity/auths"
	"test-case-majoo/entity/responses"
)

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// BP : Create user vip
func (s *Service) GetUserByID(userID string) (auths.User, *responses.Response) {
	return s.repo.GetUserByID(userID)
}

// BP : Create user vip
func (s *Service) GetMonthlyReport(userID, Month, limit, page string) (auths.MonthlyReports, *responses.Response) {
	return s.repo.GetMonthlyReport(userID, Month, limit, page)
}
