package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > 24 || b > 60 {
		return
	}
	if b >= 45 {
		fmt.Println(a, b-45)
	} else {
		if a == 0 {
			fmt.Println(23, b+15)
		} else {
			fmt.Println(a-1, b+15)
		}
	}
}
