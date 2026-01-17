package shared

import (
	"errors"
	"fmt"
)

// Queue is a generic fixed-capacity FIFO queue.
type Queue[T any] struct {
	data []T
	head int
	tail int
	size int
}

// NewQueue creates a new queue with the given capacity.
func NewQueue[T any](cap int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, cap),
		head: 0,
		tail: 0,
		size: 0,
	}
}

// Enqueue adds an element to the end of the queue in O(1) time.
func (q *Queue[T]) Enqueue(value T) error {
	if q.size >= len(q.data) {
		return errors.New("Enqueue to full queue")
	}

	q.data[q.tail] = value
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
	return nil
}

// Dequeue removes and returns the element at the front of the queue in O(1) time.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("Dequeue from empty queue")
	}

	val := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return val, nil
}

// Peek returns the element at the front of the queue without removing it in O(1) time.
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("Peek from empty queue")
	}
	return q.data[q.head], nil
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// Print prints the queue elements from front to back.
func (q *Queue[T]) Print() {
	for i := 0; i < q.size; i++ {
		if i > 0 {
			fmt.Print(" -> ")
		}
		idx := (q.head + i) % len(q.data)
		fmt.Print(q.data[idx])
	}
	fmt.Println()
}
