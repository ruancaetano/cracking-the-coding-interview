package shared

type Stack[T comparable] struct {
	data []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the value at the top of the stack
// Panics if the stack is empty
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Pop from empty stack")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// Peek returns the value at the top of the stack without removing it
// Panics if the stack is empty
func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("Peek from empty stack")
	}
	return s.data[len(s.data)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Print prints the stack elements from top to bottom
func (s *Stack[T]) Print() {
	for i := len(s.data) - 1; i >= 0; i-- {
		if i < len(s.data)-1 {
			print(" -> ")
		}
		print(s.data[i])
	}
	println()
}
