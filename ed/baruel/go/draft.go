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
    fmt.Println(fig)
}
