package main

import "fmt"

func main() {
	var items [6]int

	var max int

	var value int

	for i := 0; i < len(items); i++ {
		fmt.Scan(&value)
		items[i] = value
	}

	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {

			if items[i] < items[j] {
				max = items[j]

			} else if items[i] == items[j] {
				max = items[j]
			} else {

				if max < items[i] {
					max = items[i]
				} else {
					max = max
				}

			}

		}
	}
	shortHand()
	fmt.Println("max 1", max)

}

func shortHand() {
	array := [...]int{
		1, 2, 4, 44, 5, 66, 11,
	}

	var max int = array[0]

	for i := 0; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}

	fmt.Println("max 2", max)
}
