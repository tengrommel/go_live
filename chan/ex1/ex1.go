package main

import (
	"math/rand"
	"time"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {


}

func waitForTask()  {
	ch := make(chan string)

	go func() {
		p := <- ch
		fmt.Println("employee : recv'd signal :", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500))*time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------")
}

func waitForResult()  {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500))*time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	p := <-ch
	fmt.Println("manager : recv'd signal :", p)
	time.Sleep(time.Second)
	fmt.Println("--------------------------------------------------------")
}

func waitForFinished()  {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500))*time.Millisecond)
		close(ch)
		fmt.Println("employee : sent signal")
	}()

	_, wd := <- ch
	fmt.Println("manager : recv'd signal :", wd)

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

func pooling()  {
	ch := make(chan string)
	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++{
		ch <- "paper"
		fmt.Println("manager : sent signal :", w)
	}
	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

