package stack

import "errors"

type Stack[T comparable] struct {
	data []T
	top  int
}

func NewStack[T comparable](cap int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, cap),
		top:  -1,
	}
}

// Push adds an element to the top of the stack in O(1) time
func (s *Stack[T]) Push(value T) error {
	if s.top+1 >= len(s.data) {
		return errors.New("Push to full stack")
	}
	s.top++
	s.data[s.top] = value
	return nil
}

// Pop removes and returns the value at the top of the stack in O(1) time
// Panics if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Pop from empty stack")
	}
	val := s.data[s.top]
	s.top--
	return val, nil
}

// Peek returns the value at the top of the stack without removing it in O(1) time
// Panics if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Peek from empty stack")
	}
	return s.data[s.top], nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return s.top + 1
}

// Clean removes all items from the stack, making it empty.
func (s *Stack[T]) Clean() {
	s.top = -1
}

// Print prints the stack elements from top to bottom
func (s *Stack[T]) Print() {
	for i := s.top; i >= 0; i-- {
		if i < s.top {
			print(" -> ")
		}
		print(s.data[i])
	}
	println()
}
