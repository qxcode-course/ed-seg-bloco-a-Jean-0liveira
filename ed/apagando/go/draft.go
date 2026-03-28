package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n)

	fila := make([] int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&m)

	saiu := make(map[int]bool)

	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		saiu[x] = true
	}

	for _, p := range fila{
		if !saiu[p]{
			fmt.Print(p, " ")
		}
	}
	fmt.Println()
}
