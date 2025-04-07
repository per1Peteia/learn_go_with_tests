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
