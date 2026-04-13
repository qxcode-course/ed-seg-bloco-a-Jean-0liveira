package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    macho := make([] int, 0)
    femea := make([] int, 0)
    casais := 0

    for  i := 0; i < n; i++{
        var x int
        fmt.Scan(&x)
        
        if x > 0 {
            macho = append(macho, x)
        } else {
            femea = append(femea, x)
        }

        if len(macho) >= 1 && len(femea) >= 1{
            casais++

            macho  = macho[1:]
            femea = femea[1:]
        }
    }
    fmt.Println(casais)
}
