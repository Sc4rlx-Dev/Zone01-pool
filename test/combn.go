package main

import "fmt"


func PrintCombN(n int) {
    R := make([]int, n)
    for i := 0; i < n; i++ {
        R[i] = i
    }
    index := n - 1
    for {
        // for j := index; j >= 0; j-- {
        for k := R[index]; k <= 9; k++ {
            R[index] = k
            if index == n-1 {
                fmt.Println(R)
            } else {
                index++
                R[index] = R[index-1]
            }
        }
        index--
        if index < 0 {
            break
        }
        R[index]++
        for i := index + 1; i < n; i++ {
            R[i] = R[i-1] + 1
        }
    }
}



func main(){
    PrintCombN(1)
    PrintCombN(2)
    PrintCombN(3)
}