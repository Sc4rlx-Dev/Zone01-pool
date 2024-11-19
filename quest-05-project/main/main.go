// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	elems := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(piscine.BasicJoin(elems))
// }

package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}