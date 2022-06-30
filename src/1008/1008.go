package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(float64(a) / float64(b))
}
