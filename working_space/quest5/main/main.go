package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}



// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Compare("Hello!", "Hello!"))
// 	fmt.Println(piscine.Compare("Salut!", "lut!"))
// 	fmt.Println(piscine.Compare("Ola!", "Ol"))
// }