package chapter3

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/queue"
)

// RunChapter3Exercise4
// Queue via Stacks: Implement a queue using two stacks.
func RunChapter3Exercise4() {
	fmt.Println("Queue via Stacks (DoubleStackQueue) Example:")

	// Create a queue with some capacity.
	queue := queue.NewDoubleStackQueue(10)

	// First phase: enqueue initial items.
	firstBatch := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Enqueuing first batch of values:")
	for _, v := range firstBatch {
		if err := queue.Enqueue(v); err != nil {
			fmt.Printf("Enqueue(%d): error: %v\n", v, err)
		} else {
			fmt.Printf("Enqueued %d\n", v)
		}
	}

	// Dequeue half of the items from the first batch.
	fmt.Println("\nDequeuing half of the items:")
	toDequeue := len(firstBatch) / 2
	for i := 0; i < toDequeue; i++ {
		val, err := queue.Dequeue()
		if err != nil {
			fmt.Printf("Dequeue %d: error: %v\n", i+1, err)
			break
		}
		fmt.Printf("Dequeue %d: %d\n", i+1, val)
	}

	// Second phase: enqueue additional items.
	secondBatch := []int{7, 8, 9, 10}
	fmt.Println("\nEnqueuing second batch of values:")
	for _, v := range secondBatch {
		if err := queue.Enqueue(v); err != nil {
			fmt.Printf("Enqueue(%d): error: %v\n", v, err)
		} else {
			fmt.Printf("Enqueued %d\n", v)
		}
	}

	// Finally, dequeue everything that's left.
	fmt.Println("\nDequeuing all remaining items:")
	remaining := (len(firstBatch) - toDequeue) + len(secondBatch)
	for i := 0; i < remaining; i++ {
		val, err := queue.Dequeue()
		if err != nil {
			fmt.Printf("Dequeue %d: error: %v\n", i+1, err)
			break
		}
		fmt.Printf("Dequeue %d: %d\n", i+1, val)
	}
}
