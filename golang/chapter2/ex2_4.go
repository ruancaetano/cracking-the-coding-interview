package chapter2

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter2Exercise4
// Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
// before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
// to be after the elements less than x (see below). The partition element x can appear anywhere in the
// "right partition"; it does not need to appear between the left and right partitions.
// Example: Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5]
// Output: 3 -> 2 -> 1 -> 5 -> 8 -> 5 -> 10
func RunChapter2Exercise4() {
	// Build example list
	inputA := shared.NewLinkedListFromSlice([]int{3, 5, 8, 5, 10, 2, 1})
	inputB := shared.NewLinkedListFromSlice([]int{3, 5, 8, 5, 10, 2, 1})

	partitionVal := 5

	fmt.Print("Original list: ")
	inputA.Print()

	// The function to partition the list (to be implemented)
	buildPartitionWithSort(inputA.Head, partitionVal)
	fmt.Printf("Partitioned list around %d (sort): ", partitionVal)
	inputA.Print()

	newHead := buildStablePartition(inputB.Head, partitionVal)
	inputB.Head = newHead

	fmt.Printf("Partitioned list around %d (stable): ", partitionVal)
	inputB.Print()
}

// Partial bubble sort. Respect the scenarios, but change the elements order
// Time: N * N/2 -> N^2 / 2 -> N^2
// Space: O(1)
func buildPartitionWithSort(head *shared.Node[int], target int) {
	i := head
	var j *shared.Node[int]

	for {
		if i == nil {
			return
		}

		if i.Value < target {
			i = i.Next
			continue
		}

		j = i.Next
		for {
			if j == nil {
				break
			}

			if i.Value >= target && j.Value < i.Value {
				aux := i.Value
				i.Value = j.Value
				j.Value = aux
			}

			j = j.Next
		}

		i = i.Next
	}
}

// Stable partition, keeping order of the elements
// Time: N
// Space: O(1)
func buildStablePartition(head *shared.Node[int], target int) *shared.Node[int] {
	var belowTargetStart, belowTargetEnd *shared.Node[int]
	var aboveTargetStart, aboveTargetEnd *shared.Node[int]

	for head != nil {
		next := head.Next
		head.Next = nil // the connection will be redefined

		if head.Value < target {
			if belowTargetStart == nil {
				belowTargetStart = head
				belowTargetEnd = head
				head = next
				continue
			}

			belowTargetEnd.Next = head
			belowTargetEnd = head
			head = next
			continue
		}

		if aboveTargetStart == nil {
			aboveTargetStart = head
			aboveTargetEnd = head
			head = next
			continue
		}

		aboveTargetEnd.Next = head
		aboveTargetEnd = head
		head = next
	}

	// if the lsit doesn`t cotains any value below the target
	if belowTargetStart == nil {
		return aboveTargetStart
	}

	// connect the parts
	belowTargetEnd.Next = aboveTargetStart
	return belowTargetStart
}
