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

        if D == "L"{
            if x == 6 && y == 6 {
                y--
            } else {
                x--
            }
        } else if D == "R" {
            x++
        } else if D == "D" {
            y++
        } else if D == "U" {
            if x == 6 && y == 5{
                x--
            } else {
                y--
            }
        }

        fmt.Println(x, y)
    }
}
