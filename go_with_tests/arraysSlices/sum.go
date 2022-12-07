package arraysSlices

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, 0)
	for _, numbers := range numbersToSum {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		sums = append(sums, sum)
	}
	return sums

}
