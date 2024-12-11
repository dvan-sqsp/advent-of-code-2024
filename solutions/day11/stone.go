package day11

import (
	"strconv"
)

const (
	rule1 = 1
	rule3 = 2024
)

type Stone struct {
	value int
}

func NewStone(num int) *Stone {
	return &Stone{
		value: num,
	}
}

func (s *Stone) IsNumberLengthEven() bool {
	return len(s.string())%2 == 0
}

func (s *Stone) RuleTwo() (*Stone, *Stone) {
	numStr := s.string()
	s1, _ := strconv.Atoi(numStr[:len(numStr)/2])
	s2, _ := strconv.Atoi(numStr[len(numStr)/2:])

	return NewStone(s1), NewStone(s2)
}

func (s *Stone) RuleThree() *Stone {
	s.value *= rule3
	return s
}

func (s *Stone) RuleOne() *Stone {
	s.value = rule1
	return s
}

func (s *Stone) string() string {
	return strconv.Itoa(s.value)
}
