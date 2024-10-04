package util

import (
	"strings"
)

type StringBuilder struct {
	values []string
}

func NewStringBuilder() StringBuilder {
	return StringBuilder{
		values: make([]string, 0),
	}
}

func (s StringBuilder) Add(input string) StringBuilder {
	s.values = append(s.values, input)
	return s
}

func (s StringBuilder) GetLastInserted() string {
	last := len(s.values) - 1
	if last < 0 {
		return ""
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

	s.values = append(s.values[:last])
	return s
}

func (s StringBuilder) Build() string {
	return strings.Join(s.values, "")
}
