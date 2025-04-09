package main

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(v int) {
	s.values = append(s.values, v)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(s.values) - 1
	popped := s.values[idx]
	s.values = s.values[:idx]
	return popped, true
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(v string) {
	s.values = append(s.values, v)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	idx := len(s.values) - 1
	popped := s.values[idx]
	s.values = s.values[:idx]
	return popped, true
}

// with generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	idx := len(s.values) - 1
	popped := s.values[idx]
	s.values = s.values[:idx]
	return popped, true
}
