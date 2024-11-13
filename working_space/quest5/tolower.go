package piscine

func ToLower(s string) string {
	ch := []rune(s)
	i := 0
	for _, r := range s {
		if r >= 65 && r <= 90 {
			ch[i] = r + 32
		}
	i++
	}
return string(ch)
}