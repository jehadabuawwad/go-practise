package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 1, 2
}

func variadicSum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	a, b := vals()
	sum := plus(a, b)
	fmt.Printf("the values %v and %v sum is %v", a, b, sum)
	variadicSum(1, 2, 3, 4, 5)
}
