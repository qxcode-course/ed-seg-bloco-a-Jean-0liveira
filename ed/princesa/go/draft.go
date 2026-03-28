package main

import (
	"fmt"
)

func main() {
    var n, e int

    fmt.Scan(&n)
    fmt.Scan(&e)

    fila := make([] int, n)

    for i := 0; i < n; i++{
        fila[i] = i + 1
    }

    fmt.Print("[ ")
    for _, p := range fila {
        if p == e {
            fmt.Print(p, "> ")
        } else {
            fmt.Print(p, " ")
        }
    }
    fmt.Print("]")
}