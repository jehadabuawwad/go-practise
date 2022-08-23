package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map : ", m)

	v1 := m["j1"]
	fmt.Println("value 1 : ", v1)
	fmt.Println("length m :", len(m))
	delete(m, "k2")
	fmt.Println("new map : ", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}
