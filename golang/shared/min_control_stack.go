package shared

import "errors"

type MinControlStack struct {
	values    []int
	minValues *Stack[int]
	topIdx    int
	cap       int
}

func NewMinControlStack(cap int) *MinControlStack {
	return &MinControlStack{
		values:    make([]int, cap),
		minValues: NewStack[int](cap),
		topIdx:    -1,
		cap:       cap,
	}
}

func (s *MinControlStack) Push(v int) error {
	if s.topIdx == s.cap-1 {
		return errors.New("stack overflow")
	}

	s.topIdx++
	s.values[s.topIdx] = v

	s.setMinValue(v)
	return nil
}

func (s *MinControlStack) Pop() (int, error) {
	if s.topIdx == -1 {
		return 0, errors.New("stack is empty")
	}

	poppedValue := s.values[s.topIdx]
	s.topIdx--
	s.removeMinValue(poppedValue)
	return poppedValue, nil
}

func (s *MinControlStack) GetMinValue() (int, error) {
	return s.minValues.Peek()
}

func (s *MinControlStack) setMinValue(newValue int) error {
	if s.minValues.IsEmpty() {
		return s.minValues.Push(newValue)
	}

	minValue, err := s.minValues.Peek()
	if err != nil {
		return err
	}

	if newValue <= minValue {
		return s.minValues.Push(newValue)
	}

	return nil
}

func (s *MinControlStack) removeMinValue(poppedValue int) error {
	lastMinValue, err := s.minValues.Peek()
	if err != nil {
		return err
	}

	if poppedValue == lastMinValue {
		_, err := s.minValues.Pop()
		return err
	}
	return nil
}
