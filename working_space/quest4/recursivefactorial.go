package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
<<<<<<< HEAD
	if nb == 0 {
		return 1
	}
	if nb <= 20 {
		if nb == 1 {
			return nb
		}
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
=======
	if nb == 1 || nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
>>>>>>> 61b451f0accdd1716b7da455c5c20cdb9bb99820
}
