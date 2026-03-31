package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"slices"
)

func getMen(vet []int) []int {
	var man [] int
	 for _, v := range vet{
		if v > 0{
			man = append(man, v)
		}
	 }
	 return man
}

func getCalmWomen(vet []int) []int {
	var stress [] int
	for _, s := range vet{

		if s < 0 {
			abs := s

			if abs < 0 {
				abs = -abs
			}
			
			if abs < 10 {
				stress = append(stress, s)
			}
		}
	}
	return stress
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	stress := make([] int, len(vet))

	copy(stress, vet)

	slices.SortFunc(stress, func(a, b int) int {
		absA := a
		absB := b

		if a < 0 {
			absA = -a
		}
		if b < 0 {
			absB = -b
		}

		return absA - absB
	})

	return stress
}	

func reverse(vet []int) []int {
	inverter := make([] int, len(vet))

	tam := len(vet)

	for i, v := range vet{
		inverter[tam - 1 - i] = v
	}
	return inverter
}

func unique(vet []int) []int {
	vistos := make(map[int] bool)
	var unico [] int
	
	for _, v := range vet{
		if !vistos[v]{
			vistos[v] = true
			unico = append(unico, v)
		}
	}

	return unico
}

func repeated(vet []int) []int {
	vistos := make(map[int] bool)
	jaMap := make(map[int] bool)
	var repetidos [] int

	for _, v := range vet {
		if !vistos[v]{
			vistos[v] = false
			repetidos = append(repetidos, v)
		}
	}

	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

