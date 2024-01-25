package main

import "fmt"

func main() {
	var varName bool = false
	fmt.Println(varName)
	varName = true
	fmt.Println(varName)

	condition := true
	fmt.Println(condition)

	str := "This is a string"
	fmt.Println(str)

	newInt := 10
	fmt.Println(newInt)

	fmt.Printf("%s, %d\n", str, newInt)

}
