package main

import "fmt"

func sliceRange() {

	x := []int{1, 2, 3, 4, 5}

	for element := range x {
		fmt.Println(element)
	}

}
