package main

import "fmt"

func main() {
	var t, r int
	fmt.Scan(&t, &r)

	v := make([]string, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&v[i])
	}

	r %= t

	res := append(v[t-r:], v[:t-r]...)

	fmt.Print("[ ")
	for _, i := range res {
		fmt.Print(i, " ")
	}
	fmt.Println("]")
}
