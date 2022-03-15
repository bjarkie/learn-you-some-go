package sum

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
		fmt.Println("Sums", numbers)
	}

	fmt.Println("Total", sums)
	return sums
}
