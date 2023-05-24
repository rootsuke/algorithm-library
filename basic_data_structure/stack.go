package basic_data_structure

type Stack struct {
	stack []int
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() int {
	lastValue := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return lastValue
}
