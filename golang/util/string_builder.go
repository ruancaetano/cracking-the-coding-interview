package util

type StringBuilder struct {
	values []rune
}

func NewStringBuilder() StringBuilder {
	return StringBuilder{
		values: make([]rune, 0),
	}
}

func (s StringBuilder) InsertAtEnd(input ...rune) StringBuilder {
	s.values = append(s.values, input...)
	return s
}

func (s StringBuilder) InsertAtStart(input ...rune) StringBuilder {
	s.values = append(input, s.values...)
	return s
}

func (s StringBuilder) GetLastInserted() rune {
	last := len(s.values) - 1
	if last < 0 {
		return rune(0)
	}

	return s.values[last]
}

func (s StringBuilder) StringsCount() int {
	return len(s.values)
}

func (s StringBuilder) DeleteLast() StringBuilder {
	last := len(s.values) - 1
	if last < 0 {
		return s
	}

	s.values = s.values[:last]
	return s
}

func (s StringBuilder) Build() string {
	return string(s.values)
}
