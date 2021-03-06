package main

import (
	"strconv"
	"os"
	"fmt"
)

func main() {
	// Atoi converts a string to an int
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	result := sum(a, b)
	fmt.Println("The sum of %d and %d\n", a, b, result)
}

func sum(a, b int) int {
	return a + b
}