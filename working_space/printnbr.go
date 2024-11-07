package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if(n == 0){
		z01.PrintRune(rune(0))
	}
	if (n == -9223372036854775808){
		
	}
	if(n > 9){
		for (n >= 9){
			z01.PrintRune(rune(n / 10) + 48)
			z01.PrintRune(rune(n % 10) + 48)
			n = n % 10
		}
	}	
	if(n <= 9){
		z01.PrintRune(rune(n) + 48)
	}
}