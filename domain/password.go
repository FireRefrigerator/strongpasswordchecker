package domain

const (
	minLen   = 8
	maxLen   = 20
	maxTypes = 3
)

type Password struct {
	initialStr    string
	Len           int
	TypesNum      int
	RegularCounts []regularCount
}

type regularCount struct {
	initialChar  byte
	initialIndex int
	times        int
}

func NewPassword(s string) *Password {
	passWord := &Password{initialStr: s, Len: len(s)}
	typesInit(passWord)
	regularCountsInit(passWord)
	return passWord
}

func typesInit(passWord *Password) {
	if hasDigit(passWord.initialStr) {
		passWord.TypesNum++
	}
	if hasUpperChar(passWord.initialStr) {
		passWord.TypesNum++
	}
	if hasLowerChar(passWord.initialStr) {
		passWord.TypesNum++
	}
	if hasSpecialChar(passWord.initialStr) {
		passWord.TypesNum++
	}
}

func regularCountsInit(pwd *Password) {
	for i := 0; i < pwd.Len-2; i++ {
		c0 := pwd.initialStr[i]
		c1 := pwd.initialStr[i+1]
		c2 := pwd.initialStr[i+2]
		if ok, inc := isCharSequence(c0, c1, c2); ok {
			i = i + 2
			count := regularCount{initialChar: c0, initialIndex: i, times: 3}
			for j := i + 1; j < pwd.Len; j++ {
				if !isCharSequence2(pwd.initialStr[i], pwd.initialStr[j], inc) {
					break
				}
				i++
				count.times++
			}
			pwd.RegularCounts = append(pwd.RegularCounts, count)
		}
	}
}

func isCharSequence(c0, c1, c2 byte) (bool, int) {
	if c0 == c1 && c1 == c2 {
		return true, 0
	}
	if c0 == c1-1 && c1 == c2-1 {
		return true, 1
	}
	if c0 == c1+1 && c1 == c2+1 {
		return true, -1
	}
	return false, 0
}

func isCharSequence2(c0, c1 byte, inc int) bool {
	if c0 == c1 && inc == 0 {
		return true
	}
	if c0+1 == c1 && inc == 1 {
		return true
	}
	if c0-1 == c1 && inc == -1 {
		return true
	}
	return false
}

func hasDigit(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func hasUpperChar(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func hasLowerChar(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func hasSpecialChar(s string) bool {
	specialCharList := map[string]string{"!": "", "@": "", "#": "", "$": "", "%": "", "^": "", "&": "", "*": ""}
	for _, c := range s {
		if _, ok := specialCharList[string(c)]; ok {
			return true
		}
	}
	return false
}
