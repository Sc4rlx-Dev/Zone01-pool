package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {

	if(n == 0){
		z01.PrintRune(rune(0))
	}
	if (n == -9223372036854775808){
		s := '9'
		z01.PrintRune('-')
		z01.PrintRune(s)
		n = 223372036854775808
	}
	if(n < 0){
		z01.PrintRune('-')
		n = -n
	}
	if(n > 9){
		for (n >= 10){
			PrintNbr(n / 10)
			n = n % 10
		}
	}	
	if(n <= 9){
		z01.PrintRune(rune(n) + 48)
	}
}