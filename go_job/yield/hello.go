package main

import (
	"fmt"
	"time"
)

/*
FIFO,first in first out
{{},{},{}}
 */


func sample(message chan<- string)  {
	message <- "hello goroutine!1\n"
	message <- "hello goroutine!2\n"
	message <- "hello goroutine!3\n"
	message <- "hello goroutine!4\n"
}

func sample2(message chan string)  {
	time.Sleep(1 *time.Second)
	str := <-message
	str = str + "I'm goroutine!"
	message <- str
	close(message)
}

func main() {
	var message = make(chan string, 3)
	go sample(message)
	go sample2(message)
	time.Sleep(2*time.Second)
	for str := range message{
		fmt.Println(str)
	}
	fmt.Println("bye")
}
