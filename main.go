package main

import "fmt"

func main() {
	fmt.Println("Swap Without Variable")
	var a = 10
	var b = 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
