package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	fmt.Println(s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println(c)

	l := c[:2]
	fmt.Println(l)

}
