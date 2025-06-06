package piscine

// Sqrt(4)
func Sqrt(nb int) int {
	if nb < 0 || nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
