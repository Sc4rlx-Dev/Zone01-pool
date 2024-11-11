package piscine

import "fmt"

func BasicAtoi(s string) int {
	i := []byte(s)
	res := 0
	flag := 0
	for a := 0 ; a <= len(s) - 1; a++{
		if i[a] >= 128 && i[a] <= 255{		
			// i[a] = i[a] - 1		
			// fmt.Println(i[a])
			flag++
		}
		// fmt.Println("|")
		// fmt.Printf("%T")
		i[a] = i[a]
		
		res++		
	}
	fmt.Printf("flag : %d\n",flag)
	res = res - flag + 1
return res
}