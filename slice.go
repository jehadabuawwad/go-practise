package main

import "fmt"

func slice() {

	a := []int{1, 2, 3, 4, 5}

	x := a[0:2]

	n := append(x, 10)

	fmt.Println(n)
}
