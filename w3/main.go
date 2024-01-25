package main

import "fmt"

func main() {
	printName("Maz")
	fmt.Println(sum(10, 20))
	echo()
	functionName()
}

func echo() {
	fmt.Println("This string is printerd from echo function")
}

func sum(a int, b int) int {
	c := a + b
	return c
}

func printName(name string) {
	println(name)
}

func functionName() {

}
