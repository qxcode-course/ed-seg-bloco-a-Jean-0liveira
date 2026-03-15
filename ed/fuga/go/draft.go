package main

import "fmt"

func main() {

	var h, p, f, d int

	fmt.Scan(&h)
	fmt.Scan(&p)
	fmt.Scan(&f)
	fmt.Scan(&d)

	for {
		f = (f + d + 16) % 16

		if f == h {
			fmt.Println("S")
			break
		} else if f == p {
			fmt.Println("N")
			break
		}
	}
}
