package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	b1 := b / 100
	b2 := b % 100 / 10
	b3 := b % 100 % 10
	fmt.Println(a * b3)
	fmt.Println(a * b2)
	fmt.Println(a * b1)
	fmt.Println(a*b3 + 10*a*b2 + 100*a*b1)
}
