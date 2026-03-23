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

	saiu := make([] int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&saiu[i])
	}
	fmt.Println(n)
	fmt.Println(fila)
	fmt.Println(m)
	fmt.Println(saiu)
}
