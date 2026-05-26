package arrays_and_slices

func Sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	result := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}

	return result
}
