package shared

import (
	"errors"
	"time"
)

type Pet struct {
	name      string
	kind      string
	entryDate time.Time
}

func NewPet(name, kind string, entryDate time.Time) (Pet, error) {
	if kind != "dog" && kind != "cat" {
		return Pet{}, errors.New("invalid pet kind")
	}

	return Pet{
		name:      name,
		kind:      kind,
		entryDate: entryDate,
	}, nil
}

func (p Pet) Name() string {
	return p.name
}

func (p Pet) Kind() string {
	return p.kind
}

func (p Pet) EntryDate() time.Time {
	return p.entryDate
}
