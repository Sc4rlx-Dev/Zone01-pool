package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	var del string = ", "

	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			z01.PrintRune(rune(i/10 + 48))
			z01.PrintRune(rune(i%10 + 48))
			z01.PrintRune(rune(del[1]))
			z01.PrintRune(rune(j/10 + 48))
			z01.PrintRune(rune(j%10 + 48))
			if(i == 98){
				break
			}
			z01.PrintRune(rune(del[0]))
			z01.PrintRune(rune(del[1]))
		}
	}
	
z01.PrintRune(rune('\n'))
}