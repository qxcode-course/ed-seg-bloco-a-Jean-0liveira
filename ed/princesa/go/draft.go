package main

import (
	"fmt"
)

func main() {
    var n, e int

    fmt.Scan(&n)
    fmt.Scan(&e)

    vivos := make([] bool, n)

    for v := range vivos{
        vivos[v] = true
    }

    pos := e - 1
    num_vivos := n

    for {

        fmt.Print("[ ")

        for i := 0; i < len(vivos); i++{

            if vivos[i] {
                if i == pos {
                    fmt.Print(i+1, "> ")
                } else {
                    fmt.Print(i+1, " ")
                }
            }
        }

        fmt.Println("]")

        if num_vivos == 1{
            break
        }

        matar := pos

        for {
            matar = (matar + 1) % n

            if vivos[matar]{
                break
            }
        }
    
        vivos[matar] = false
        num_vivos--

        pos = matar

        for {
            pos = (pos + 1) % n

            if vivos[pos]{
                break
            }
        }
    }
}