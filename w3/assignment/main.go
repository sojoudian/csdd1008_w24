package main

import "fmt"

func main() {
	fmt.Println(multiply(2, 3))
}

// Function two
// this function will calculate the multiplication of two numbers with int type
// and return the result with int type
func multiply(a int, b int) int {
	return a * b
}
