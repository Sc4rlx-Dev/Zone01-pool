package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb <= 0 {
		if( nb == 0){
			return 1
		}
		return 0
	}

	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
