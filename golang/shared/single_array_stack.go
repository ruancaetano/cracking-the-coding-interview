package shared

import (
	"errors"
	"fmt"
)

type SingleArrayStack struct {
	partitionCap  int
	partitionRefs []int
	values        []int
}

func NewSingleArrayStack(cap, partitions int) *SingleArrayStack {
	if cap < 0 {
		panic("capacity must be positive")
	}

	if partitions < 0 {
		panic("partitions must be positive")
	}

	if cap%partitions != 0 {
		panic("capacity must be divisible by partitions")
	}

	values := make([]int, cap)
	partitionRefs := make([]int, partitions)

	for i := range partitionRefs {
		partitionRefs[i] = -1
	}

	return &SingleArrayStack{
		partitionRefs: partitionRefs,
		values:        values,
		partitionCap:  cap / partitions,
	}
}

func (s *SingleArrayStack) Push(p, v int) error {
	if !s.isValidPartition(p) {
		return errors.New("partition out of range")
	}

	if s.isPartitionFull(p) {
		return errors.New("partition is full")
	}

	if s.isPartitionEmpty(p) {
		s.partitionRefs[p] = s.getFirstPartitionIndex(p)
	} else {
		s.partitionRefs[p]++
	}

	s.values[s.partitionRefs[p]] = v
	return nil
}

func (s *SingleArrayStack) Pop(p int) (int, error) {
	if !s.isValidPartition(p) {
		return 0, errors.New("partition out of range")
	}

	if s.isPartitionEmpty(p) {
		return 0, errors.New("partition is empty")
	}

	poppedValue := s.values[s.partitionRefs[p]]

	if s.partitionRefs[p] == s.getFirstPartitionIndex(p) {
		s.partitionRefs[p] = -1
	} else {
		s.partitionRefs[p]--
	}

	return poppedValue, nil
}

func (s *SingleArrayStack) Peek(p int) (int, error) {
	if !s.isValidPartition(p) {
		return 0, errors.New("partition out of range")
	}

	if s.isPartitionEmpty(p) {
		return 0, errors.New("partition is empty")
	}

	return s.values[s.partitionRefs[p]], nil
}

func (s *SingleArrayStack) Print() {
	fmt.Println("Values: ", s.values)
	fmt.Println("Partition Refs: ", s.partitionRefs)

	for p, ref := range s.partitionRefs {
		fmt.Printf("Stack %d: ", p+1)
		firstIndex := s.getFirstPartitionIndex(p)

		for i := firstIndex; i <= ref; i++ {
			fmt.Printf("%d ", s.values[i])
		}
		fmt.Println()
	}
}

func (s *SingleArrayStack) isValidPartition(p int) bool {
	return p >= 0 && p < len(s.partitionRefs)
}

func (s *SingleArrayStack) isPartitionEmpty(p int) bool {
	return s.partitionRefs[p] == -1
}

func (s *SingleArrayStack) isPartitionFull(p int) bool {
	return s.partitionRefs[p] == s.getLastPartitionIndex(p)
}

func (s *SingleArrayStack) getFirstPartitionIndex(p int) int {
	return p * s.partitionCap
}

func (s *SingleArrayStack) getLastPartitionIndex(p int) int {
	return ((p + 1) * s.partitionCap) - 1
}
