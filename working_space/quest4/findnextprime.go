package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	for {
		flag := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			return nb
		}
		nb++
	}
}

// func FindNextPrime(nb int) int {

// 	if nb <= 1 {
// 		return 2
// 	}
// 	for {
// 		if isPrime(nb) {
// 			return nb
// 		}
// 		nb++
// 	}
// }

// func isPrime(n int) bool {
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
