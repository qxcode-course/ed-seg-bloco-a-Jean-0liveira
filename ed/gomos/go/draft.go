package main

import "fmt"

func main() {
    var Q int
    var D string

    fmt.Scan(&Q)
    fmt.Scan(&D)
    
    for i := 0; i < Q; i++ {
        var x int
        var y int 

        fmt.Scan(&x)
        fmt.Scan(&y)

        

        fmt.Println(x, y)
    }
}
