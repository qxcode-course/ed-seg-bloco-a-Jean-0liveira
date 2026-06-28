package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	input := scanner.Text()

	var antes []rune
	var depois []rune

	for _, c := range input {
		switch c {
		case 'D':
			depois = depois[1:]
		case '>':
			if len(depois) > 0 {
				antes = append(antes, depois[0])
				depois = depois[1:]
			}
		case 'B':
			if len(antes) > 0 {
				antes = antes[:len(antes)-1]
			}
		case '<':
			if len(antes) > 0 {
				index := len(antes) - 1
				depois = append([]rune{antes[index]}, depois...)
				antes = antes[:index]
			}
		case 'R':
			antes = append(antes, '\n')
		default:
			antes = append(antes, c)
		}
	}
	fmt.Print(string(antes))
	fmt.Print("|")
	fmt.Println(string(depois))
}
