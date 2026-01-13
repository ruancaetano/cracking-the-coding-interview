package shared

import (
	"errors"
	"fmt"
)

type SetOfStacks struct {
	stacks [][]int
	top    int
	cap    int
}

func NewSetOfStacks(cap int) *SetOfStacks {
	return &SetOfStacks{
		stacks: make([][]int, 0, cap),
		top:    -1,
		cap:    cap,
	}
}

func (s *SetOfStacks) Push(v int) {
	if len(s.stacks) == 0 {
		s.stacks = append(s.stacks, make([]int, s.cap))
	}

	stackIdx := len(s.stacks) - 1
	s.top += 1

	if s.top == s.cap {
		stackIdx += 1
		s.stacks = append(s.stacks, make([]int, s.cap))
		s.top = 0
	}

	s.stacks[stackIdx][s.top] = v
}

func (s *SetOfStacks) Pop() (int, error) {
	if len(s.stacks) == 0 {
		return 0, errors.New("stack is empty")
	}

	stackIdx := len(s.stacks) - 1
	poppedValue := s.stacks[stackIdx][s.top]
	s.top--

	// last stack is not empty?
	if s.top > -1 {
		return poppedValue, nil
	}

	// Remove the last stack when it's empty
	s.stacks = s.stacks[:stackIdx]
	stackIdx--

	// reset top after removing last stack
	if stackIdx == -1 {
		s.top = -1
	} else {
		s.top = len(s.stacks[stackIdx]) - 1
	}

	return poppedValue, nil
}

func (s *SetOfStacks) Print() {
	if len(s.stacks) == 0 {
		fmt.Println("Stack is empty")
		return
	}

	for si, stack := range s.stacks {
		fmt.Printf("Stack %d: ", si)
		limit := s.cap
		// for last stack, limit to (top+1) elements
		if si == len(s.stacks)-1 {
			limit = (s.top + 1)
		}
		fmt.Print("[")
		for i := 0; i < limit && i < len(stack); i++ {
			fmt.Print(stack[i])
			if i < limit-1 && i < len(stack)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	}
}
