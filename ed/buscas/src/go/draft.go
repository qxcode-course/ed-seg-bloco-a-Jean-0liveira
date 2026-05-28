package main

import "fmt"

func main() {
	var num_entrada int
	fmt.Scan(&num_entrada)

	frequencia := make(map[string]int)
	for i := 0; i < num_entrada; i++ {
		var str string
		fmt.Scan(&str)
		frequencia[str]++
	}

	var num_consulta int
	fmt.Scan(&num_consulta)

	consulta := make([]string, num_consulta)
	for i := 0; i < num_consulta; i++ {
		fmt.Scan(&consulta[i])
	}

	for i, c := range consulta {
		fmt.Printf("%d", frequencia[c])
		if i < num_consulta-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
