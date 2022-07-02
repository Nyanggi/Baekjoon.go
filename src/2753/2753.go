package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	if checkIfLeapYear(a) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

func checkIfLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
