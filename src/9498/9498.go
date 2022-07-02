package main

import "fmt"

func main() {
	var score int
	fmt.Scanln(&score)
	if score >= 90 || score == 100 {
		fmt.Println("A")
	}
	if score >= 80 && score <= 89 {
		fmt.Println("B")
	}
	if score >= 70 && score <= 79 {
		fmt.Println("C")
	}
	if score >= 60 && score <= 69 {
		fmt.Println("D")
	}
	if score < 60 {
		fmt.Println("F")
	}
}
