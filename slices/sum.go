package main

import "slices"

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		var fakeZero []int
		if len(numbers) == 0 {
			fakeZero = slices.Concat([]int{0}, numbers)
		} else {
			fakeZero = numbers
		}

		result = append(result, Sum(fakeZero[1:]))
	}

	return result
}
