package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	m := len(grid)
	n := len(grid[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == word[0] {
				if backtrack(grid, r, c, 0, word) {
					return true
				}
			}
		}
	}

	return false
}

func backtrack(grid [][]byte, r int, c int, index int, word string) bool {

	if index == len(word) {
		return true
	}

	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}

	palavra := grid[r][c]

	grid[r][c] = '#'

	busca := backtrack(grid, r+1, c, index+1, word) ||
		backtrack(grid, r-1, c, index+1, word) ||
		backtrack(grid, r, c+1, index+1, word) ||
		backtrack(grid, r, c-1, index+1, word)

	grid[r][c] = palavra

	return busca
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
