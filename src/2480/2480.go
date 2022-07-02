package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sum := 0
	if a == b && b == c && c == a {
		fmt.Println(10000 + a*1000)
	} else if a == b || a == c {
		fmt.Println(1000 + a*100)
	} else if b == c {
		fmt.Println(1000 + b*100)
	} else {
		z := math.Max(float64(a), float64(b))
		z = math.Max(z, float64(c))
		sum = int(z) * 100
		fmt.Println(sum)
	}
}
