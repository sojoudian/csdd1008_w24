package main

import "fmt"

func main() {
	fmt.Println(echo("Function one, created by Maziar"))
	fmt.Println(multiply(2, 3))
	newPrint("Function three, created by NExt person")
}

// Function one
// created by maziar
// this function will echo the input string
func echo(s string) string {
	return s
}

// Function two
// this function will calculate the multiplication of two numbers with int type
// and return the result with int type
func multiply(a int, b int) int {
	return a * b
}

// Function three
// created by next person
// this function will print the input string to the command line
func newPrint(s string) {
	fmt.Println(s)
}
