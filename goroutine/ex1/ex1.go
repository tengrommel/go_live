package main

import (
	"fmt"
	"sync"
)

func init()  {
	// Allocate
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function
	go func() {
		lowercase()
		wg.Done()
	}()

	// Create a goroutine from the uppercase function
	go func() {
		uppercase()
		wg.Done()
	}()

	// Wait for the goroutines to finish
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func uppercase() {

	// Display the alphabet three times
	for count:=0;count<3;count++ {
		for r := 'a'; r <= 'z'; r++{
			fmt.Printf("%c ", r)
		}
	}
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {

	// Display the alphabet three times
	for count:=0;count<3;count++ {
		for r := 'A'; r <= 'Z'; r++{
			fmt.Printf("%c ", r)
		}
	}
}
