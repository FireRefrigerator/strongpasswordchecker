package domain

type PasswordModifyStep struct {
	*Password
	Step int
}

func (s *PasswordModifyStep) increaseLen() {
	s.Len++
}

func (s *PasswordModifyStep) decreaseLen() {
	s.Len--
}

func (s *PasswordModifyStep) increaseTypesNum() {
	if s.TypesNum < maxTypes {
		s.TypesNum++
	}
}

func (s *PasswordModifyStep) increaseStep() {
	s.Step++
}

func (s *PasswordModifyStep) consumerRegularCounts(quota int) {
	if len(s.RegularCounts) != 0 {
		s.RegularCounts[0].times -= quota
		if s.RegularCounts[0].times < 3 {
			s.RegularCounts = s.RegularCounts[1:]
		}
		return
	}
}

func (s *PasswordModifyStep) consumerRegularCountsByPrio() {
	if len(s.RegularCounts) == 0 {
		return
	}

	// 用3取模， 生成map[int]int, key为0，1，2的子集合， value为对应元素的角标
	m := map[int]int{}
	for i := 0; i < len(s.RegularCounts); i++ {
		o := s.RegularCounts[i].times % 3
		m[o] = i
	}

	// 取模值最小的元素的角标
	expectIndex := -1
	for j := 0; j < 3; j++ {
		if v, ok := m[j]; ok {
			expectIndex = v
			break
		}
	}

	// 删除一个字符进行consumer操作
	s.RegularCounts[expectIndex].times--
	if s.RegularCounts[expectIndex].times < 3 {
		s.RegularCounts = append(s.RegularCounts[:expectIndex],
			s.RegularCounts[expectIndex+1:]...)
	}
}
