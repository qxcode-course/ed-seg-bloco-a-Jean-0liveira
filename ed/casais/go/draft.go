package main

import (
	"fmt"
)

func main() {
    
    var n int

    fmt.Scan(&n)

    macho := make([] int, 0)
    femea := make([] int, 0)
    casais := 0
 
    for i := 0; i < n; i++ {
        var j int
        fmt.Scan(&j)

        if j > 0 {
            macho = append(macho, j)
        } else {
            femea = append(femea, j)
        }

        if len(macho) >= 1 && len(femea) >= 1{
            casais++

            macho = macho[1:]
            femea = femea[1:]
        }
    }

    fmt.Println(casais)
}
