package chapter3

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/stack"
)

// RunChapter3Exercise1
// Three in One: Describe how you could use a single array to implement three stacks.

// ----------------------------------------------------------------------------------
// ANSWER
//
// We can implement three stacks in a single array by dividing the array
// into three fixed partitions of equal size.
//
// For example, given an array of size 15, we create three partitions of size 5:
//
// indexes  0–4   - stack 1
// indexes  5–9   - stack 2
// indexes 10–14  - stack 3
//
// Each stack maintains its own top pointer that always stays within its partition.
// The trade-off is that space is inflexible: small stacks still reserve the same
// capacity as large ones.
//
// To improve space usage, we could design a more dynamic structure where each stack
// has its own (start, size, capacity) metadata and, when one stack needs more
// space, we shift neighboring stacks within the array to free room.

func RunChapter3Exercise1() {

	fmt.Printf("========== Creating stack ==========\n\n")
	stack := stack.NewSingleArrayStack(15, 3)
	stack.Print()

	fmt.Printf("\n\n========== Pushing values ==========\n\n")

	stack.Push(0, 1)
	stack.Push(0, 2)
	stack.Push(0, 3)
	stack.Push(1, 4)
	stack.Push(1, 5)
	stack.Push(1, 6)
	stack.Push(2, 7)
	stack.Push(2, 8)
	stack.Push(2, 9)
	stack.Print()

	fmt.Printf("\n\n========== Popping values ==========\n\n")

	stack.Pop(0)
	stack.Pop(1)
	stack.Pop(2)
	stack.Print()

	fmt.Printf("\n\n========== Peeking values ==========\n\n")

	value, _ := stack.Peek(0)
	fmt.Println("Peeked value: ", value)
	value, _ = stack.Peek(1)
	fmt.Println("Peeked value: ", value)
	value, _ = stack.Peek(2)
	fmt.Println("Peeked value: ", value)
}
