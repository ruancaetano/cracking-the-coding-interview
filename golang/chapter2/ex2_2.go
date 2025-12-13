package chapter2

import (
	"errors"
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter2Exercise2
// Input: A linked list and a value k
// The task: Return the kth to last element of the linked list.
// Example input: 1 -> 2 -> 3 -> 4 -> 5, k = 2
// Expected output: 4
func RunChapter2Exercise2() {
	input := shared.NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})
	k := 2

	input.Print()
	result, err := getKthLastElement(input.Head, k)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Println("Last Kth (", k, "): ", result)
}

// Two pointer/runner strategy
// Time: O(K + N - K) -> worse case K is equal N -> O(N - 0) -> O(N)
// Space: O(1)
func getKthLastElement(head *shared.Node[int], k int) (int, error) {
	if k <= 0 {
		return 0, errors.New("k must be positive")
	}

	firstPointer := head
	secondPointer := head

	for i := 1; i <= k; i++ {
		if secondPointer == nil {
			return 0, errors.New("invalid array")
		}
		secondPointer = secondPointer.Next
	}

	for {
		if secondPointer == nil {
			break
		}
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next
	}

	return firstPointer.Value, nil
}
