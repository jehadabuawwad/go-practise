package main

import "fmt"

func flow() {

	num := 22

	if num == 22 {
		for x := 0; x < 5; x++ {
			switch x {
			case 0:
				fmt.Println(" ")
				fmt.Println("ok")
			case 1:
				fmt.Println("not-ok")
			}
		}
	}

}
