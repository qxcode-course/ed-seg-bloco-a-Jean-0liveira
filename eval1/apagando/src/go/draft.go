package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fila := make([] int, n)
	
	for i := 0; i < n; i++{
		fmt.Scan(&fila[i])
	}

	var m int

	fmt.Scan(&m)

	saiu := make(map[int] bool)

	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		saiu[x] = true
	}

	for _, f := range fila{
		if !saiu[f] {
			fmt.Print(f, " ")
		}
	}
	fmt.Println()
}
