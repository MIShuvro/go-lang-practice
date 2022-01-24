package main

import "fmt"

func main() {
	var items [7]int

	var search int

	var value int
	var frist = 0
	var last = len(items) - 1
	var middle = (frist + last) / 2
	fmt.Println("Enter a Value of Search: ")
	fmt.Scan(&search)
	for i := 0; i < len(items); i++ {
		fmt.Scan(&value)
		items[i] = value
	}

	for frist <= last {
		if items[middle] < search {
			frist = middle + 1
			middle = (frist + last) / 2
		} else if items[middle] == search {
			fmt.Println("Found the position at: ", middle)
			break
		} else {
			last = middle - 1
			middle = (frist + last) / 2
		}
	}

	if frist > last {
		fmt.Println("Not Found")
	}
}
