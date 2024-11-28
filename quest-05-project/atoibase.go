package piscine

import "fmt"
func AtoiBase(s string, base string) int {
	result := 0
	sign := 1
	if base[0] == '-' || base[0] == '+' {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	digit := map[byte]int{}
	for i := 0; i < len(base); i++ {
		digit[base[i]] = i
	}
	fmt.Println(digit)
	for i := 0; i < len(s); i++ {
		digit := digit[s[i]]
		result = result*len(base) + digit
	}
	return result * sign
}
