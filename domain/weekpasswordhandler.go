package domain

type Handler interface {
	Match(pwd *Password) bool
	Action(pwdStep *PasswordModifyStep)
}

/* 密码长度小于8的处理器 */
type LessThanHandler struct {
}

var Repeat = func(handler Handler, pwdStep *PasswordModifyStep) *PasswordModifyStep {
	for {
		if !handler.Match(pwdStep.Password) {
			break
		}
		handler.Action(pwdStep)
	}
	return pwdStep
}

func (h *LessThanHandler) Match(pwd *Password) bool {
	return pwd.Len < minLen
}

func (h *LessThanHandler) Action(pwdStep *PasswordModifyStep) {
	pwdStep.increaseLen()
	pwdStep.increaseTypesNum()
	pwdStep.consumerRegularCounts(2)
	pwdStep.increaseStep()
}

/* 密码长度大于20的处理器 */
type MoreThanHandler struct {
}

func (h *MoreThanHandler) Match(pwd *Password) bool {
	return pwd.Len > maxLen
}

func (h *MoreThanHandler) Action(pwdStep *PasswordModifyStep) {
	pwdStep.decreaseLen()
	pwdStep.consumerRegularCountsByPrio()
	pwdStep.increaseStep()
}

/* 密码存在3个字符序列的处理器 */
type SequenceHandler struct {
}

func (h *SequenceHandler) Match(pwd *Password) bool {
	return len(pwd.RegularCounts) > 0
}

func (h *SequenceHandler) Action(pwdStep *PasswordModifyStep) {
	pwdStep.increaseTypesNum()
	pwdStep.consumerRegularCounts(3)
	pwdStep.increaseStep()
}

/* 密码类型小于3的处理器 */
type TypesHandler struct {
}

func (h *TypesHandler) Match(pwd *Password) bool {
	return pwd.TypesNum < 3
}

func (h *TypesHandler) Action(pwdStep *PasswordModifyStep) {
	pwdStep.increaseTypesNum()
	pwdStep.consumerRegularCounts(3)
	pwdStep.increaseStep()
}

/* top 10弱密码处理器 */
type Top10Handler struct {
}

func (h *Top10Handler) Match(pwd *Password) bool {
	return inBlackList(pwd.initialStr)
}

func (h *Top10Handler) Action(pwdStep *PasswordModifyStep) {
}

func inBlackList(pwd string) bool {
	blackList := pwdBlackList()
	if _, ok := blackList[pwd]; ok {
		return true
	}
	return false
}

func pwdBlackList() map[string]string {
	return map[string]string{
		"123456":     "",
		"abcdef":     "",
		"654321":     "",
		"qazwsx":     "",
		"123!@#":     "",
		"qwertasdfg": "",
		"11111":      "",
		"22222":      "",
		"33333":      "",
		"66666":      "",
	}
}
