package main

import "github.com/01-edu/z01"

func main(){
	for i := 0; i <= 9 ; i++{
		z01.PrintRune(rune(i + '0'))
	}
	z01.PrintRune('\n')
}