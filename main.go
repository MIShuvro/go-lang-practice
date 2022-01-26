package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort")
	var items [5]int
	var value int
	for i := 0; i < len(items); i++ {
		fmt.Scan(&value)
		items[i] = value
	}
	for a := 0; a < len(items)-1; a++ {
		for b := 0; b < len(items)-1; b++ {
			if items[b] > items[b+1] {
				var temp int
				temp = items[b]
				items[b] = items[b+1]
				items[b+1] = temp
			}
		}
	}
	fmt.Println(items)
	var search int
	fmt.Println("Enter search value:")
	fmt.Scan(&search)
	var frist = 0
	var last = len(items) - 1
	var middle = (frist + last) / 2
	for j := 0; j < len(items); j++ {

		if items[middle] < search {
			frist = middle + 1
			middle = (frist + last) / 2
		} else if items[middle] == search {
			fmt.Print("Position", middle)
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
