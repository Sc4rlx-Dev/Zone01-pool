package piscine

import "fmt"

func QuadAT(x,y int) {
	tmpx := x
	tmpy := x
	flag := 0

	//handle the negative numbers
	if(x <= 0|| y <= 0 ){
		return 
	}
	if(x >= 0 && y >= 0){
		fmt.Print("o")
		x -= 2
		for(x > 0 ){
			fmt.Print("-")
			x--
		}
		fmt.Print("o")
		fmt.Print("\n")
		y -= 2
		for(y > 0){
			fmt.Print("|")
			y--
			flag = 1
			if(flag == 1){
				tmpx -= 2
				for(tmpx > 0){
					fmt.Print(" ")
					tmpx--
				}
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		fmt.Print("o")
		if(flag == 1){
			tmpy -= 2	
			for(tmpy > 0 ){
				fmt.Print("-")
				tmpy--
			}
		}
	}
	fmt.Print("o")
	// fmt.Print("o")
}