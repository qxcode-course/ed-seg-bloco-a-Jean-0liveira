package main

import (
	"bufio"
	"fmt"
	"os"
)

func busca(grid [][]byte, l int, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] == '0' {
		return
	}

	grid[l][c] = '0'

	busca(grid, l+1, c)
	busca(grid, l-1, c)
	busca(grid, l, c+1)
	busca(grid, l, c-1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	nl := len(grid)
	nc := len(grid[0])

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				count++
				busca(grid, i, j)
			}
		}
	}
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
