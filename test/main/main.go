package main

import (
	"os"
	// "fmt"
	"github.com/01-edu/z01"
)

func Test(args []string) []string {
	v1 := map[string]string{
		"str4": args[4],
		"str3": args[3],
		"str2": args[2],
		"str1": args[1],
	}
	var result []string
	for _, value := range v1 {
		var s string
		for _, r := range value {
				s += string(r)
		}
		result = append(result, s)
	}
	// fmt.Println(result)
	return result
}


func main() {
	args := Test(os.Args)
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
