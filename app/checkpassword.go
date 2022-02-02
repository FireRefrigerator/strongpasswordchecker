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
	repeat(&domain.LessThanHandler{}, passwordStep)
	repeat(&domain.MoreThanHandler{}, passwordStep)
	repeat(&domain.SequenceHandler{}, passwordStep)
	repeat(&domain.TypesHandler{}, passwordStep)
	repeat(&domain.Top10Handler{}, passwordStep)
	return passwordStep.Step
}

var repeat = func(handler domain.Handler, pwdStep *domain.PasswordModifyStep) *domain.PasswordModifyStep {
	for {
		if !handler.Match(pwdStep.Password) {
			break
		}
		handler.Action(pwdStep)
	}
	return pwdStep
}
