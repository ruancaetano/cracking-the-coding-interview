package stack

type SortStack struct {
	stack    *Stack[int]
	auxStack *Stack[int]
}

func NewSortStack(cap int) *SortStack {
	return &SortStack{
		stack:    NewStack[int](cap),
		auxStack: NewStack[int](cap),
	}
}

func (s *SortStack) Push(v int) error {
	return s.stack.Push(v)
}

func (s *SortStack) Pop() (int, error) {
	return s.stack.Pop()
}

func (s *SortStack) Peek() (int, error) {
	return s.stack.Peek()
}

func (s *SortStack) Sort() error {
	if s.stack.IsEmpty() {
		return nil
	}

	for !s.stack.IsEmpty() {
		poppedValue, _ := s.stack.Pop()

		for {
			if s.auxStack.IsEmpty() {
				s.auxStack.Push(poppedValue)
				break
			}

			auxTopValue, _ := s.auxStack.Peek()
			if poppedValue > auxTopValue {
				s.auxStack.Push(poppedValue)
				break
			}

			auxPoppedValue, _ := s.auxStack.Pop()
			s.stack.Push(auxPoppedValue)
		}
	}

	for !s.auxStack.IsEmpty() {
		poppedValue, _ := s.auxStack.Pop()
		s.stack.Push(poppedValue)
	}

	return nil
}

func (s *SortStack) Print() {
	s.stack.Print()
}
