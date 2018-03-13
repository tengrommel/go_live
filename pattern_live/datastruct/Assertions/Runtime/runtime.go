package main

import (
	"math/rand"
	"time"
	"fmt"
)

type car struct {}

func (car)String() string {
	return "Vroom!"
}

type cloud struct {}

func (cloud)String() string {
	return "Big Data!"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Create a slice of the Stringer interface values.
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	// Let's run this experiment ten times.
	for i:=0;i<10;i++ {
		rn := rand.Intn(2)
		if v, is := mvs[rn].(cloud);is{
			fmt.Println("Got lucky:", v)
			continue
		}
		fmt.Println("Got Unlucky")
	}
}
