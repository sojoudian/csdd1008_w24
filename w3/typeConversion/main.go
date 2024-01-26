package main

import "fmt"

func main() {
	pi := 3.14
	newInt := int(pi)
	fmt.Println(newInt)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", newInt)
}
