package piscine

func LastRune(s string) rune {
	i := 0
	for range s{
		i++
	}
	return []rune(s)[i - 1]
}