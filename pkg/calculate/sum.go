package calculate

func SumArray(values ...int) int {
	result := 0

	for _, value := range values {
		result += value
	}

	return result
}
