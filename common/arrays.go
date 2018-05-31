package common

func Reverse(numbers []int) []int {
	midpoint := len(numbers) / 2
	arraylen := len(numbers)
	for i := 0; i < midpoint; i++ {
		j := arraylen - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
