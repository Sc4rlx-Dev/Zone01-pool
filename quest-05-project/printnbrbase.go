package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		for _, r := range "-9223372036854775808" {
			z01.PrintRune(r)
		}
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	var result string
	length := len(base)

	for nbr > 0 {
		digit := nbr % length
		result = string(base[digit]) + result
		nbr /= length
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
}