package main

import "fmt"

func main() {
    
    var n int

    fmt.Scan(&n)

    a := make([] int, n)
    b := make([] int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&a[i], &b[i])
    }

    menorValor := 101
    vencedor := -1

    for i := 0; i < n; i++ {
        a := a[i]
        b := b[i]
    }

}
