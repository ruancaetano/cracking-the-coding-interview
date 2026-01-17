package chapter3

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter3Exercise5
// Sort Stack: Write a program to sort a stack such that the smallest items are on the top, using at most one additional stack.
func RunChapter3Exercise5() {
	fmt.Println("Chapter 3 - Exercise 5: Sort Stack")

	sortStack := shared.NewSortStack(10)

	// push some unsorted values
	values := []int{3, 5, 1, 4, 2}
	for _, v := range values {
		if err := sortStack.Push(v); err != nil {
			fmt.Println("error pushing value:", err)
			return
		}
	}

	fmt.Print("Original stack (top -> bottom): ")
	sortStack.Print()

	// sort the stack
	if err := sortStack.Sort(); err != nil {
		fmt.Println("error sorting stack:", err)
		return
	}

	fmt.Print("Sorted stack   (top -> bottom): ")
	sortStack.Print()
}
