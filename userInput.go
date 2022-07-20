package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func userInput() {
	fmt.Println("Type you born year: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Your Age : %d", 2022-input)
}
