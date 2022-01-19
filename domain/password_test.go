package domain

import "testing"

func TestNewPassword(t *testing.T) {
	password := NewPassword("123ABCddd***")
	if password.initialStr != "123ABCddd***" {
		t.Errorf("test new password err, initialStr shold be %s actual %s", "123ABCddd***", password.initialStr)
	}
	if password.Len != 12 {
		t.Errorf("test new password err, Len shold be %d actual %d", 12, password.Len)
	}
	if password.TypesNum != 4 {
		t.Errorf("test new password err, TypesNum shold be %d actual %d", 4, password.TypesNum)
	}
	if len(password.RegularCounts) != 4 {
		t.Errorf("test new password err, len RegularCounts shold be %d actual %d", 4, len(password.RegularCounts))
	} 

	password = NewPassword("123ABCddd**&")
	if password.initialStr != "123ABCddd**&" {
		t.Errorf("test new password err, initialStr shold be %s actual %s", "123ABCddd**&", password.initialStr)
	}
	if password.Len != 12 {
		t.Errorf("test new password err, Len shold be %d actual %d", 12, password.Len)
	}
	if password.TypesNum != 4 {
		t.Errorf("test new password err, TypesNum shold be %d actual %d", 4, password.TypesNum)
	}
	if len(password.RegularCounts) != 3 {
		t.Errorf("test new password err, len RegularCounts shold be %d actual %d", 3, len(password.RegularCounts))
	}
}
