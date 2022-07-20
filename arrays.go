package main

import "fmt"

func arrays() {

	arr := [3]int{1, 2, 3}
	fmt.Println(len(arr))

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	println(sum)
}
