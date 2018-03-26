package main

import (
	"strconv"
	"fmt"
	"time"
)


func go_sample1(ch chan string)  {
	for i:=0;i<19;i++{
		ch <- "I'm sample1 num:" + strconv.Itoa(i)
		time.Sleep(1*time.Second)
	}
}

func go_sample2(ch chan int)  {
	for i:=0;i<10;i++{
		ch <- i
		time.Sleep(2*time.Second)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i:=0;i<10;i++ {
		go go_sample1(ch1)
		go go_sample2(ch2)
	}

	for i:=0;i<1000;i++  {
		select {
		case str, ch1Check := <- ch1:
			if !ch1Check {
				fmt.Println("ch1 check failed")
			}
			fmt.Println(str)
		case p, ch2Check := <- ch2:
			if !ch2Check {
				fmt.Println("ch2 check failed")
			}
			fmt.Println(p)
		}
	}
	time.Sleep(60*time.Second)
	fmt.Println("end!")
}
