package main

import (
	"fmt"
)

func main() {
    
    var n int

    fmt.Scan(&n)

    macho := make([]int, 0)
    femea := make([] int, 0)

    for i := 0; i < n; i++ {
        fmt.Scan(&n)

        if n > 0 {
            macho = append(macho, n)
        } else {
            femea = append(femea, n)
        }
    }

    fmt.Println(n)
}
