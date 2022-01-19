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
