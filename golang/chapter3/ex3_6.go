package chapter3

import (
	"fmt"
	"time"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter3Exercise6
// Animal Shelter: Implement a data structure to manage a shelter that holds only dogs and cats and operates on a strictly FIFO basis.
func RunChapter3Exercise6() {
	fmt.Println("Chapter 3 - Exercise 6: Animal Shelter")

	shelter := shared.NewPetShelter(10)
	now := time.Now()

	// create some pets with increasing entry dates
	petsData := []struct {
		name   string
		kind   string
		offset time.Duration
	}{

		{"Buddy", "dog", 0},
		{"Mittens", "cat", time.Second},
		{"Rex", "dog", 2 * time.Second},
		{"Luna", "cat", 3 * time.Second},
	}

	for _, pd := range petsData {
		pet, err := shared.NewPet(pd.name, pd.kind, now.Add(pd.offset))
		if err != nil {
			fmt.Println("error creating pet:", err)
			return
		}
		if err := shelter.Enqueue(pet); err != nil {
			fmt.Println("error enqueuing pet:", err)
			return
		}
	}

	// helper to print results of dequeues
	dequeueAny := func() {
		pet, err := shelter.DequeueAny()
		if err != nil {
			fmt.Println("DequeueAny error:", err)
			return
		}
		fmt.Printf("DequeueAny -> %s (%s)\n", pet.Name(), pet.Kind())
	}

	dequeueDog := func() {
		pet, err := shelter.DequeueDog()
		if err != nil {
			fmt.Println("DequeueDog error:", err)
			return
		}
		fmt.Printf("DequeueDog -> %s (%s)\n", pet.Name(), pet.Kind())
	}

	dequeueCat := func() {
		pet, err := shelter.DequeueCat()
		if err != nil {
			fmt.Println("DequeueCat error:", err)
			return
		}
		fmt.Printf("DequeueCat -> %s (%s)\n", pet.Name(), pet.Kind())
	}

	// demonstrate behavior
	dequeueAny()
	dequeueDog()
	dequeueCat()
	dequeueAny()
}
