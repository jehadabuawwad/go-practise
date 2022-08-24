package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 10
	return &p

}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(&person{name: "Hani", age: 20})
	fmt.Println(newPerson("Jon"))

}
