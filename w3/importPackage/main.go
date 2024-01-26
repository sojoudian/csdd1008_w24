package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(3, 16))
	fmt.Println(Echo("from echo function"))
}

func sum(a int, b int) int {
	c := a + b
	return c
}

func Echo(a string) string {
	return a
}
