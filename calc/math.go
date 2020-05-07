package calc

// Add returns sum of slice of integers
func Add(numbers ...iint) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
