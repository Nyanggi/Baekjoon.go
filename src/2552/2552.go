package main

import "fmt"

func main() {
	var h, m, c int
	fmt.Scanln(&h, &m)
	fmt.Scanln(&c)
	m = m + c
	for m >= 60 {
		m = m - 60
		h += 1

		if h >= 24 {
			h = 0
		}
	}
	fmt.Println(h, m)
}
