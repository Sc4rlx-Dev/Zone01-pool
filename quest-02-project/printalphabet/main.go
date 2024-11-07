package main 

import "github.com/01-edu/z01"

func main(){

	ch := 97
	for ch <= 122{
		z01.PrintRune(rune (ch))
		ch++
	}
	z01.PrintRune('\n')
}