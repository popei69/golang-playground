package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, slice := range numbersToSum {
		result = append(result, Sum(slice))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, slice := range numbersToSum {
		if len(slice) <= 1 {
			result = append(result, 0)
			continue
		}
		tail := slice[1:]
		result = append(result, Sum(tail))
	}

	return result
}
