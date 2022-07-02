package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(checkQuadrant(a, b))
}

func checkQuadrant(a, b int) int {
	if a > 0 && b > 0 {
		return 1
	} else if a < 0 && b > 0 {
		return 2
	} else if a < 0 && b < 0 {
		return 3
	} else if a > 0 && b < 0 {
		return 4
	} else {
		return 0
	}
}
