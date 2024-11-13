package piscine

func Concat(str1 string, str2 string) string {
	concat := ""
	for i := 0 ; i < len(str1) ;i++ {
		concat += string(str1[i])
	}
	for i := 0; i < len(str2) ; i++ {
		concat += string(str2[i])
	}
	return concat
}
