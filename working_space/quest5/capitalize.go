package piscine 
//hello! world How are You
// output : Hello! World How Are You 

func Capitalize(s string) string {
	r := []rune(s)
	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] -= 32
	}
	for i := 1 ; i < len(s) - 1 ; i++ {
		if r[i] == ' ' && (r[i + 1] >= 'a' && r[i + 1] <= 'z') {
			r[i + 1] -= 32
		} else if r[i] != ' ' && (r[i + 1] >= 'A' && r[i + 1] <= 'Z') {
			r[i + 1] += 32
		}
	}
	return string(r)
}


// func Capitalize(s string) string {
// 	r := []rune(s)
// 	if r[0] >= 'a' && r[0] <= 'z' {
// 		r[0] -= 32
// 	}

// 	for i := 1; i < len(r); i++ {
// 		if r[i-1] == ' ' && r[i] >= 'a' && r[i] <= 'z' {
// 			r[i] -= 32
// 		} else if r[i-1] != ' ' && r[i] >= 'A' && r[i] <= 'Z' {
// 			r[i] += 32
// 		}
// 	}
// return string(r)
// }