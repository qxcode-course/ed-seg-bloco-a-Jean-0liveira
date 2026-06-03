package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Fprintf(&result, ", %d", s.data[i])
	}
	result.WriteString("]")
	return result.String()
}

func (s *Set) AcharPosicao(value int) (int, bool) {
	low := 0
	high := s.size - 1

	for low <= high {
		mid := low + (high-low)/2
		if s.data[mid] == value {
			return mid, true
		} else if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low, false
}

func (s *Set) insert(value int) {
	idx, found := s.AcharPosicao(value)

	if found {
		return
	}

	if s.size == s.capacity {
		newCap := 1
		if s.capacity > 0 {
			newCap = s.capacity * 2
		}
		s.Reserve(newCap)
	}

	for i := s.size; i > idx; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[idx] = value
	s.size++
}

func (s *Set) Reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) Contains(value int) bool {
	_, found := s.AcharPosicao(value)
	return found
}

func (s *Set) Erase(value int) bool {
	idx, found := s.AcharPosicao(value)
	if !found {
		return false
	}

	for i := idx; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.size--
	return true
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
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
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.insert(value)
			}
		case "show":
			fmt.Println(s.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !s.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if s.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println(false)
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
