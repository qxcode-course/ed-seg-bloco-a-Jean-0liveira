package main

import "fmt"
import "math"

func main() {
    
    var n int

    fmt.Scan(&n)

    pa := make([] int, n)
    pb := make([] int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&pa[i], &pb[i])
    }

    menorValor := 101
    vencedor := -1

    for i := 0; i < n; i++ {
        
        a := pa[i]
        b := pb[i]

        if a < 10 || b < 10{
            continue
        }

        dif := int(math.Abs(float64(a - b)))

        if dif < menorValor {
            menorValor = dif
            vencedor = i
        }
    }
    if vencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor)
    }
}
