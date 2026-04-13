package main

import "fmt"

func main() {
    var qtd_total_fig int
    var qtd_baruel_fig int
    
    fmt.Scan(&qtd_total_fig)
    fmt.Scan(&qtd_baruel_fig)

    figs := make([] int, qtd_baruel_fig)

    for i:= 0; i < qtd_baruel_fig; i++{
        fmt.Scan(&figs[i])
    }

    fmt.Print(figs)

    reps := make(map[int]int)

    for _, f := range figs{
        reps[f]++
    }
}
