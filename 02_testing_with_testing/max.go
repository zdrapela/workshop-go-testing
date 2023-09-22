package testing_with_testing

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
