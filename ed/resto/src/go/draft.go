package main

import "fmt"

func divisao(n int){
    if n == 0{
        return
    }

    prox_num := n / 2
    resto := n % 2

    divisao(prox_num)

    fmt.Println(prox_num, resto)
}

func main() {
    var num int

    fmt.Scan(&num)

    divisao(num)
}
