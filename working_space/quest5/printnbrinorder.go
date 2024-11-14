package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var digits [20]int
	nbr := 0

	for n > 0 {
		digits[nbr] = n % 10
		n /= 10
		nbr++
	}

	for i := 0 ; i < nbr - 1 ; i++ {
		for j := 0 ; j < nbr - i - 1 ; j++ {
			if digits[j] > digits[j + 1] {
				digits[j] , digits[j + 1] = digits[j + 1] , digits[j]
			}
		}
	}

	for i:= 0 ; i < nbr ; i++ {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
