package piscine
// package main

import "piscine"


func printcomb(){
	var del string = ", "
	for i:= 0; i <= 7; i++{
		for j:= i + 1; j <= 8; j++{
			for k:= j + 1; k <= 9; k++{
				z01.PrintRune(rune(i) + 48)
				z01.PrintRune(rune(j) + 48)
				z01.PrintRune(rune(k) + 48)
					if(i == 7){
						break						
					}
					z01.PrintRune(rune(del[0]))
					z01.PrintRune(rune(del[1]))
			}
		}
	}
z01.PrintRune('\n')
}

// func main(){
// 	printcomb()
// }