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

    reps := make(map[int]int)

    for _, f := range figs{
        reps[f]++
    }
     
    pri_rep := true
    repetiu := false

    for _, f := range figs{
        if reps[f] > 1{
            if !pri_rep{
                fmt.Print(" ")
            }
            fmt.Print(f)
            repetiu = true
            pri_rep = false
            reps[f]--
        }
    }

    if !repetiu {
        fmt.Println("N")
    } else {
        fmt.Println()
    }

    pri_fal := true
    faltou := false

    for i := 1; i <= qtd_total_fig; i++{
        if reps[i] == 0{
            if !pri_fal{
                fmt.Print(" ")
            }
            fmt.Print(i)
            pri_fal = false
            faltou = true
        }
    }
    if faltou{
        fmt.Println()
    } else {
        fmt.Println("N")
    }
}
