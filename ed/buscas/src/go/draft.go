package main

import "fmt"

func main() {
	var num_consulta int
	fmt.Scan(&num_consulta)

	consulta := make([]string, num_consulta)
	for i := 0; i < num_consulta; i++ {
		fmt.Scan(&consulta[i])
	}

	var num_entrada int
	fmt.Scan(&num_entrada)

	frequencia := make(map[string]int)
	for i := 0; i < num_entrada; i++ {
		var str string
		fmt.Scan(&str)
		frequencia[str]++
	}
	for _, c := range consulta {
		fmt.Printf("%d", frequencia[c])
	}
}
