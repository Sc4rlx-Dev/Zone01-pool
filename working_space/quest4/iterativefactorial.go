package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return 1
	}

	for i := 1; i <= nb; i++ {
		result *= i
		if result < 0 || result > 9223372036854775807 {
			return 0
		}
	}
	return result
}
