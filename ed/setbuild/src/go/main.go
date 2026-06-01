package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) Show() string {
	if s.size == 0 {
		return "[]"
	}
	var result strings.Builder
	fmt.Fprintf(&result, "[%d", s.data[0])
	for i := 1; i < s.size; i++ {
		fmt.Fprintf(&result, " %d", s.data[i])
	}
	result.WriteString("]")
	return result.String()
}

func (s *Set) AcharPosicao(value int) (int, bool) {
	low := 0
	high := s.size - 1

	for low <= high {
		mid := low + (high-low)/2
		if s.data
	}
}

func (s *Set) insert(value int) {

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			// value, _ := strconv.Atoi(parts[1])
			// v = NewSet(value)
		case "insert":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
			fmt.Println(v.Show())
		case "erase":
			// value, _ := strconv.Atoi(parts[1])
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
