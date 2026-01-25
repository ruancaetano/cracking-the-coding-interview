package chapter3

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/stack"
)

// RunChapter3Exercise2
// Stack Min: Design a stack that supports push, pop, and retrieving the minimum element in O(1) time.
func RunChapter3Exercise2() {
	fmt.Println("Stack Min (MinControlStack) Example:")

	stack := stack.NewMinControlStack(10)
	vals := []int{5, 6, 3, 7, 2, 2, 8}

	fmt.Println("Pushing values onto stack:")
	for _, v := range vals {
		err := stack.Push(v)
		if err != nil {
			fmt.Printf("Push(%d): error: %v\n", v, err)
		} else {
			min, _ := stack.GetMinValue()
			fmt.Printf("Pushed %d, current min: %d\n", v, min)
		}
	}

	fmt.Println("\nPopping values and printing min after each pop:")
	for i := 0; i < len(vals); i++ {
		val, err := stack.Pop()
		if err != nil {
			fmt.Printf("Pop: error: %v\n", err)
			break
		}
		fmt.Printf("Popped %d", val)
		min, minErr := stack.GetMinValue()
		if minErr != nil {
			fmt.Printf(", stack is now empty.\n")
		} else {
			fmt.Printf(", new min: %d\n", min)
		}
	}
}
