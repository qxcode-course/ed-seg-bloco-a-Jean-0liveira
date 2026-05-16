package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	s := "["

	for i, v := range vet {
		if i > 0 {
			s += ", "
		}
		s += strconv.Itoa(v)
	}
	s += "]"

	return s
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	s := "["

	for i := len(vet) - 1; i >= 0; i-- {
		if i < len(vet)-1 {
			s += ", "
		}
		s += strconv.Itoa(vet[i])
	}
	s += "]"
	return s
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	var rec func(int, int)

	rec = func(i, j int) {
		if i >= j {
			return
		}

		vet[i], vet[j] = vet[j], vet[i]

		rec(i+1, j-1)
	}
	rec(0, len(vet)-1)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var rec func(int) int

	rec = func(i int) int {
		if i >= len(vet) {
			return 0
		}
		return vet[i] + rec(i+1)
	}

	return rec(0)
}

func mult(vet []int) int {
	var rec func(int) int

	rec = func(i int) int {
		if i >= len(vet) {
			return 1
		}
		return vet[i] * rec(i+1)
	}
	return rec(0)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(int) int

	rec = func(i int) int {
		if i == len(vet)-1 {
			return i
		}

		menor := rec(i + 1)

		if vet[i] <= vet[menor] {
			return i
		}
		return menor
	}
	return rec(0)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
