package piscine

func TrimAtoi(s string) int {
	nbr := 0
	sign := 1 
	foundSign := false
	foundNbr := false

	for _, ch := range s {
		if ch == '-' && !foundNbr && !foundSign {
			sign = -1
			foundSign = true
		} else if ch >= '0' && ch <= '9' {
			foundNbr = true
			nbr = nbr * 10 + int(ch - '0')
		}
	}
return sign * nbr
}