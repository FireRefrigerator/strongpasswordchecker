package app

import (
	"github.com/strongpasswordchecker/domain"
)

func CheckPassword(str string) bool {
	password := domain.NewPassword(str)
	if password.Len < 8 || password.Len > 20 {
		return false
	}
	if password.TypesNum < 3 {
		return false
	}
	if len(password.RegularCounts) > 0 {
		return false
	}
	return true
}

func CheckPasswordStep(str string) int {
	password := domain.NewPassword(str)
	passwordStep := &domain.PasswordModifyStep{Password: password}
	domain.Repeat(&domain.LessThanHandler{}, passwordStep)
	domain.Repeat(&domain.MoreThanHandler{}, passwordStep)
	domain.Repeat(&domain.SequenceHandler{}, passwordStep)
	domain.Repeat(&domain.TypesHandler{}, passwordStep)
	domain.Repeat(&domain.Top10Handler{}, passwordStep)
	return passwordStep.Step
}
