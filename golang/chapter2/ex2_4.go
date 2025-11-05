package chapter2

// RunChapter2Exercise4
// Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
// before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
// to be after the elements less than x (see below). The partition element x can appear anywhere in the
// "right partition"; it does not need to appear between the left and right partitions.
// Example: Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5]
// Output: 3 -> 2 -> 1 -> 5 -> 8 -> 5 -> 10
import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

func RunChapter2Exercise4() {
	// Build example list
	input := util.NewLinkedListFromSlice([]int{3, 5, 8, 5, 10, 2, 1})
	partitionVal := 5

	fmt.Print("Original list: ")
	input.Print()

	// The function to partition the list (to be implemented)
	buildPartition(input.Head, partitionVal)

	fmt.Printf("Partitioned list around %d: ", partitionVal)
	input.Print()
}

func buildPartition(head *util.Node[int], target int) {
}
