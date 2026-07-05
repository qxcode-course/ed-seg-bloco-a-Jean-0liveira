package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Soma(index int, somaAtual int, alvo int, vSoma []int) bool {

	if somaAtual == alvo {
		return true
	}

	if somaAtual > alvo || index >= len(vSoma) {
		return false
	}

	if Soma(index+1, somaAtual, alvo, vSoma) {
		return true
	}

	if Soma(index+1, somaAtual+vSoma[index], alvo, vSoma) {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	line := strings.Fields(scanner.Text())

	if !scanner.Scan() {
		return
	}

	secondLine := strings.Fields(scanner.Text())

	n, _ := strconv.Atoi(line[0])
	alvo, _ := strconv.Atoi(line[1])

	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i], _ = strconv.Atoi(secondLine[i])
	}

	if Soma(0, 0, alvo, v) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
