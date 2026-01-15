package shared

import "errors"

type DoubleStackQueue struct {
	inputStack  *Stack[int]
	outputStack *Stack[int]
	cap         int
}

func NewDoubleStackQueue(cap int) *DoubleStackQueue {
	return &DoubleStackQueue{
		inputStack:  NewStack[int](cap),
		outputStack: NewStack[int](cap),
		cap:         cap,
	}
}

func (d *DoubleStackQueue) Enqueue(v int) error {
	if d.inputStack.Size() == d.cap {
		return errors.New("queue is full")
	}

	d.inputStack.Push(v)
	return nil
}

func (d *DoubleStackQueue) Dequeue() (int, error) {
	if d.inputStack.IsEmpty() && d.outputStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	if d.outputStack.IsEmpty() {
		if err := d.fromInputToOutputStack(); err != nil {
			return 0, err
		}
	}

	return d.outputStack.Pop()
}

func (d *DoubleStackQueue) fromInputToOutputStack() error {
	for !d.inputStack.IsEmpty() {
		popped, err := d.inputStack.Pop()
		if err != nil {
			return err
		}

		if err := d.outputStack.Push(popped); err != nil {
			return err
		}
	}

	return nil
}
