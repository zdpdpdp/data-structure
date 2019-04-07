package stack

import "github.com/pkg/errors"

var emptyError = errors.New("stack is empty")

type Stack struct {
	container []int
}

func NewStack() *Stack {
	return &Stack{
		container: make([]int, 0),
	}
}

func (s *Stack) Pop() (int, error) {
	l := len(s.container)

	if l == 0 {
		return 0, emptyError
	}
	res := s.container[l-1]

	s.container = s.container[:l-1]
	return res, nil
}

func (s *Stack) Push(data int) {
	s.container = append(s.container, data)
}

func (s *Stack) len() int {
	return len(s.container)
}
