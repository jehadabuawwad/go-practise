package main

import "fmt"

func mAp() {
	hello()

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	fmt.Println(mp)

}
