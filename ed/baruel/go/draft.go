package main

import "fmt"

func main() {
    var qtd_fig_album int
    var qtd_fig_baruel int 

    fmt.Scan(&qtd_fig_album)
    fmt.Scan(&qtd_fig_baruel)

    fig := make([] int, qtd_fig_baruel)
    
    for i := 0; i < qtd_fig_baruel; i++ {
        fmt.Scan(&fig[i])
    }

    cont := make(map[int]int)

    for _, f := range fig{
        cont[f]++
    }

    repetiu := false
    primeiro_rep := true

    for _, f := range fig{
        if cont[f] > 1 {
            if !primeiro_rep {
                fmt.Print(" ")
            }
            fmt.Print(f)
            cont[f]--
            primeiro_rep = false
            repetiu = true
        }
    }

    if repetiu{
        fmt.Println()
    } else {
        fmt.Println("N")
    }

    faltou := false
    primeiro_fal := true

    for i := 1; i <= qtd_fig_album; i++ {
        if cont[i] == 0{
            if !primeiro_fal{
                fmt.Print(" ")
            }
            fmt.Print(i)
            faltou = true
            primeiro_fal = false
        }
    }

    if faltou{
        fmt.Println()
    } else {
        fmt.Println("N")
    }
}
