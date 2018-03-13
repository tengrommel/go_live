package main

import "fmt"

func main() {
	greet("Jane", "Done")
}

func greet(fname string, lname string)  {
	fmt.Println(fname, lname)
}
