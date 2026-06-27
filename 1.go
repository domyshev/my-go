package main

import "fmt"

func main() {
	var a string = "this is a string"
	var b int = 555

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(fmt.Sprintf("This is a concatination of string and integer: %s %d", a, b))
	fmt.Println(fmt.Sprint("Or like this: ", a, " ", b))
}
