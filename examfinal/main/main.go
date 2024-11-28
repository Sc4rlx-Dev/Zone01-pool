// package main

// import (
// 	"os"
// 	// "fmt"
// 	"github.com/01-edu/z01"
// )

// func Test(args []string) []string {
// 	v1 := map[string]string{
// 		"str4": args[4],
// 		"str3": args[3],
// 		"str2": args[2],
// 		"str1": args[1],
// 	}
// 	var result []string
// 	for _, value := range v1 {
// 		var s string
// 		for _, r := range value {
// 				s += string(r)
// 		}
// 		result = append(result, s)
// 	}
// 	// fmt.Println(result)
// 	return result
// }
// package main

// import (
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		os.Stdout.Write([]byte{'\n'}) // Write a newline character
// 		return
// 	}
// 	arg1 := os.Args[1]
// 	arg2 := os.Args[2]
// 	seen := make([]bool, 255) // Array to track printed characters

// 	for _, char := range arg1 { // Iterate over each character in the first string
// 		if !seen[char] { // If the character hasn't been printed yet
// 			seen[char] = true
// 			os.Stdout.Write([]byte{byte(char)}) // Write the character to stdout
// 		}
// 	}

// 	for _, char := range arg2 { // Iterate over each character in the second string
// 		if !seen[char] { // If the character hasn't been printed yet
// 			seen[char] = true
// 			os.Stdout.Write([]byte{byte(char)}) // Write the character to stdout
// 		}
// 	}

// 	os.Stdout.Write([]byte{'\n'}) // Write a newline character
// }
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	args := os.Args[1:]
// 	union := args[0] + args[1]
// 	result := []rune{}
// 	if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
// 		fmt.Println()
// 		return
// 	}
// 	for _, char := range union { // range over union
// 		found := false // use found so we don't add similar characters
// 		for _, n := range result {
// 			if char == n {
// 				found = true
// 				break
// 			}
// 		}
// 		if found == false {
// 			result = append(result, char)
// 		}
// 	}
// 	fmt.Println(string(result))
// }

// checkpoint
package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		seen := make([]bool,255) 
		for i := 1; i <= 2; i++ { 
			arg := os.Args[i]
			for j := 0; j < len(arg); j++ {
				if !seen[arg[j]] { 
					seen[arg[j]] = true
					os.Stdout.Write([]byte{arg[j]}) 
				}
			}
		}
	}
	os.Stdout.Write([]byte{'\n'}) 
}


