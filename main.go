package main

import "fmt"

func main() {
	var items [5]int
	var a int
	var search_value int
	for i := 0; i < len(items); i++ {
		fmt.Scan(&a)
		items[i] = a
	}

	fmt.Print("Please Enter a Search Value:")
	fmt.Scan(&search_value)

	for j := 0; j < len(items); j++ {
		if items[j] == search_value {
			fmt.Println("Found The Value at position: ", j)
			fmt.Println(1)

		}
	}
}
