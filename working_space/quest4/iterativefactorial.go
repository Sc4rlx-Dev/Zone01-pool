package piscine

<<<<<<< HEAD
func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return 1
	}

	for i := 1; i <= nb; i++ {
		result *= i
		if result < 0 || result > 9223372036854775807 {
			return 0
		}
	}
	return result
}
=======
import "fmt"

func Iterativefactorial(nb int) int {

	result := 1
	if nb <= 0{
		return 0
	}

	for i := 1; i <= nb; i++{
		result = result * i

		if(result < 0 || result > 9223372036854775807){
			return 0
		}
	}
	fmt.Printf("result:   %d\n", result)
	return result
}
>>>>>>> 61b451f0accdd1716b7da455c5c20cdb9bb99820
