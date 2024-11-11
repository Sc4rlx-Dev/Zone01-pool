package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
    if x <= 0 || y <= 0 {
        return
    }

    xMax := x
    yMax := y

    for i := 1; i <= yMax; i++ {
        for j := 1; j <= xMax; j++ {
            if i == 1 && j == 1 {
                z01.PrintRune('/')
            } else if i == 1 && j == xMax {
                z01.PrintRune('\\')
            } else if i == yMax && j == 1 {
                z01.PrintRune('\\')
            } else if i == yMax && j == xMax {
                z01.PrintRune('/')
            } else if i == 1 || i == yMax {
                z01.PrintRune('*')
            } else if j == 1 || j == xMax {
                z01.PrintRune('*')
            } else {
                z01.PrintRune(' ')
            }
        }
        z01.PrintRune('\n')
    }
}
