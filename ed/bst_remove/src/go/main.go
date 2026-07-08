package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insertNode(root *Node, v int) *Node {
	if root == nil {
		return &Node{Value: v}
	}
	if v < root.Value {
		root.Left = insertNode(root.Left, v)
	} else if v > root.Value {
		root.Right = insertNode(root.Right, v)
	}
	return root
}

func BstInsert(values []int) *Node {
	// TODO
	var root *Node
	for _, v := range values {
		root = insertNode(root, v)
	}
	return root
}

func maior(Node *Node) *Node {
	arvore := Node
	for arvore.Right != nil {
		arvore = arvore.Right
	}
	return arvore
}

func BstRemove(node *Node, value int) *Node {
	// TODO
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = BstRemove(node.Left, value)
	} else if value > node.Value {
		node.Right = BstRemove(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		prox := maior(node.Left)
		node.Value = prox.Value
		node.Left = BstRemove(node.Left, prox.Value)
	}

	return node
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	scanner.Scan()
	toRemove, _ := strconv.Atoi(scanner.Text())

	_ = toRemove // Ignora o valor a ser removido, pois não está implementado
	root := BstInsert(values)
	fmt.Println("original:")
	BShow(root, "") // Chama a função de impressão formatada
	root = BstRemove(root, toRemove)
	fmt.Println("modificado:")
	BShow(root, "") // Chama a função de impressão formatada da árvore modificada
}
