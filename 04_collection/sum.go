package main

func Sum(numbers []int) int {
	sum := func(acc, x int) int { return acc + x }
	return Reduce(numbers, sum, 0)
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}

func Reduce[T any](iterable []T, f func(T, T) T, initialValue T) T {
	var result = initialValue
	for _, element := range iterable {
		result = f(result, element)
	}
	return result
}
