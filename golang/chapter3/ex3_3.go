package chapter3

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/stack"
)

// RunChapter3Exercise3
// Stack of Plates: Implement a data structure SetOfStacks that behaves like a single stack but is composed of multiple stacks.
func RunChapter3Exercise3() {
	fmt.Printf("Example: SetOfStacks with sub-stack capacity 3\n\n")

	set := stack.NewSetOfStacks(3)

	fmt.Printf("\n=== Fill Stack ===\n\n")
	for i := 1; i <= 10; i++ {
		set.Push(i)
		fmt.Printf("Pushed %d\n", i)
	}

	fmt.Printf("\n=== Stacks After Push ===\n\n")
	set.Print()

	fmt.Printf("\n=== Popping elements ===\n\n")

	for i := 0; i < 10; i++ {
		val, err := set.Pop()
		if err != nil {
			fmt.Printf("Pop %d: %v\n", i+1, err)
			break
		} else {
			fmt.Printf("Pop %d: %d\n", i+1, val)
		}

		set.Print()
		fmt.Println()
	}
}
