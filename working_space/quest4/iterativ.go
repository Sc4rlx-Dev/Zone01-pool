package piscine

import "fmt"

func Iterativefactorial(index int) int {

	result := 1
	if index <= 0{
		return 0
	}

	for i := 1; i <= index; i++{
		result = result * i

		if(result < 0 || result > 9223372036854775807){
			return 0
		}
		// fmt.Printf("\n")
		fmt.Printf("i     :   %d\n", i)
		// fmt.Printf("\n")
		fmt.Printf("index :   %d", index)
	}
	fmt.Printf("result:   %d\n", result)
	return result
}