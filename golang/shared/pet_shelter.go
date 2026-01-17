package shared

type PetShelter struct {
	dogs *Queue[Pet]
	cats *Queue[Pet]
}

// NewShelter creates a new Shelter with queues for dogs and cats.
func NewPetShelter(cap int) *PetShelter {
	return &PetShelter{
		dogs: NewQueue[Pet](cap),
		cats: NewQueue[Pet](cap),
	}
}

func (s *PetShelter) Enqueue(pet Pet) error {
	if pet.kind == "dog" {
		return s.dogs.Enqueue(pet)
	}
	return s.cats.Enqueue(pet)
}

func (s *PetShelter) DequeueDog() (Pet, error) {
	return s.dogs.Dequeue()
}

func (s *PetShelter) DequeueCat() (Pet, error) {
	return s.cats.Dequeue()
}

func (s *PetShelter) DequeueAny() (Pet, error) {
	if s.cats.IsEmpty() {
		return s.dogs.Dequeue()
	}

	if s.dogs.IsEmpty() {
		return s.cats.Dequeue()
	}

	dog, _ := s.dogs.Peek()
	cat, _ := s.cats.Peek()

	if dog.EntryDate().Before(cat.EntryDate()) {
		return s.dogs.Dequeue()
	}
	return s.cats.Dequeue()
}
