package sum

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
		// fmt.Println("Loop through array", number)
	}
	return sum
}
