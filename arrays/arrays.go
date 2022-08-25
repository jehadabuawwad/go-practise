package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	a[2] = 10

	fmt.Println(a)
	fmt.Println(a[2])

	var x [2][5]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			x[i][j] = i + j
		}
	}
	fmt.Println(x)

}
