package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["11"]= 1
	m["22"]= 2
	fmt.Println(m)

	delete(m, "11")

	for k, v := range m{
		fmt.Println(k)
		fmt.Println(v)
	}
}
