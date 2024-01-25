package main

import "fmt"

func main() {
	printName("Maz")
	fmt.Println(sum(10, 20))
	echo()
	functionName()
	secondEcho("This is from secondEcho function")
	thirdEcho("This is from thirdEcho function")
	//fmt.Println(fourthEcho("This is from fourthEcho function"))
	a := fourthEcho("This is from fourthEcho function")
	fmt.Println(a)
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

func secondEcho(s string) {
	fmt.Println(s)
}

func thirdEcho(s string) string {
	fmt.Println(s)
	return s
}

func fourthEcho(s string) string {
	return s
}
