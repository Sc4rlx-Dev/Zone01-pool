package piscine

// func StrRev(s string) string {
// 	cv := []rune(s)
// 	for i, j := 0, len(cv)-1; i < j; i, j = i+1, j-1 {
// 		cv[i],cv[j] = cv[j],cv[i]
// 	}
// 	return string(cv)
// }
func StrRev(s string) string {
	rev := ""
	i := 0
	j := 0
	for range s{
		j++
	}
	for i = j - 1; i >= 0 ; i--{
		rev += string(s[i])
	}
	return rev
}


