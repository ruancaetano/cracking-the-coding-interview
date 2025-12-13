package chapter2

import "github.com/ruancaetano/cracking-the-coding-interview/golang/shared"

// RunChapter2Exercise1
// Remove Dups: Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
func RunChapter2Exercise1() {
	// Use a more comprehensive test case with multiple duplicates and scattered order
	input := shared.NewLinkedListFromSlice([]int{2, 3, 5, 3, 7, 2, 9, 5, 11, 3, 2, 11})
	inputTwo := shared.NewLinkedListFromSlice([]int{2, 3, 5, 3, 7, 2, 9, 5, 11, 3, 2, 11})
	removeDuplicatesWithHashMap(input.Head)
	removeDuplicatesWithoutHashMap(inputTwo.Head)
	// expected result: 2 -> 3 -> 5 -> 7 -> 9 -> 11
	input.Print()
	inputTwo.Print()
}

// Solve using a buffer
// Time: O(N)
// Space: O(N)
// Where N is the list length
func removeDuplicatesWithHashMap(head *shared.Node[int]) {
	found := map[int]bool{}

	current := head
	var previous *shared.Node[int]

	for {
		if current == nil {
			break
		}

		if found[current.Value] {
			// remove current node
			previous.Next = current.Next
			current = current.Next
			continue
		}

		found[current.Value] = true
		previous = current
		current = current.Next
	}
}

// Solve without a buffer
// Time: O(N^2)
// Space: O(1)
// Where N is the list length
func removeDuplicatesWithoutHashMap(head *shared.Node[int]) {
	i := head

	for {
		if i == nil {
			break
		}

		jPrevious := i
		j := i.Next
		for {
			if j == nil {
				break
			}

			if i.Value == j.Value {
				// remove j
				jPrevious.Next = j.Next
				j = j.Next
				continue
			}

			jPrevious = j
			j = j.Next
		}

		i = i.Next
	}
}
