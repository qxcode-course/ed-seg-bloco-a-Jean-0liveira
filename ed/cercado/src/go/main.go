package main

import (
	"bufio"
	"fmt"
	"os"
)

func busca(board [][]byte, l int, c int) {
	if l < 0 || c < 0 || l >= len(board) || c >= len(board[0]) || board[l][c] != 'O' {
		return
	}

	board[l][c] = '#'

	busca(board, l+1, c)
	busca(board, l-1, c)
	busca(board, l, c+1)
	busca(board, l, c-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	nl := len(board)
	nc := len(board[0])

	for i := 0; i < nl; i++ {
		if board[i][0] == 'O' {
			busca(board, i, 0)
		}
		if board[i][nc-1] == 'O' {
			busca(board, i, nc-1)
		}
	}

	for j := 0; j < nc; j++ {
		if board[0][j] == 'O' {
			busca(board, 0, j)
		}
		if board[nl-1][j] == 'O' {
			busca(board, nl-1, j)
		}
	}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
