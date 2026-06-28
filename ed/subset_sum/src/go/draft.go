package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Soma() bool {
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
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i], _ = strconv.Atoi(secondLine[i])
	}

	fmt.Println(Soma())
}
