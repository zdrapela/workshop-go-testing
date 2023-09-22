package main

import "fmt"

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func TestMax(numbers []int, expected int) string {
	output := "Pass"
	actual := Max(numbers)
	if actual != expected {
		output = fmt.Sprintf("Expected %v, but instead got %v", expected, actual)
	}
	return output
}

func main() {
	fmt.Println(TestMax([]int{1, 2, 3, 4}, 4))
	// fmt.Println(TestMax([]int{4, 2, 1, 4}, 3))
	// fmt.Println(TestMax([]int{0, 0, 0, 0}, 1))
}
