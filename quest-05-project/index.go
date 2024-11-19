package piscine

func Index(s string, toFind string) int {

for i := 0; i <= len(s)-len(toFind); i++ {
    match := false
    for j := 0; j < len(toFind); j++ {
        if s[i+j] == toFind[j] {
            match = true
            break
        }
    }
    if match {
        return i
    }
}
return -1
// for i := 0; i <= len(s)-len(toFind); i++ {
//         if s[i:i+len(toFind)] == toFind {
//             return i
//         }
//     }
// return -1
}