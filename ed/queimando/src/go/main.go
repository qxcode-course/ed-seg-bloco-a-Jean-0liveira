package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	L int
	C int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()

	stack.Push(Pos{L: l, C: c})

	nl := len(grid)
	if nl == 0 {
		return
	}

	nc := len(grid[0])

	for !stack.IsEmpty() {
		retirar := stack.Pop()

		if retirar.L < 0 || retirar.C < 0 || retirar.L >= nl || retirar.C >= nc {
			continue
		}

		if grid[retirar.L][retirar.C] == '#' {
			grid[retirar.L][retirar.C] = 'o'

			stack.Push(Pos{L: retirar.L + 1, C: retirar.C})
			stack.Push(Pos{L: retirar.L - 1, C: retirar.C})
			stack.Push(Pos{L: retirar.L, C: retirar.C + 1})
			stack.Push(Pos{L: retirar.L, C: retirar.C - 1})
		}
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
