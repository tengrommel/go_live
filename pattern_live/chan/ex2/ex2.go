package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create an unbuffered channel.
	court := make(chan int)
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)
	// Launch two players.
	go func() {
		player("Venus", court)
		wg.Done()
	}()
	go func() {
		player("Alice", court)
		wg.Done()
	}()
	court <- 1
	wg.Wait()
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int)  {
	for  {
		// Wait for the ball to be hit back to us.
		ball, wd := <- court
		if !wd {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// Pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13==0{
			fmt.Printf("Player %s Missed\n", name)
			// Close the channel to signal we lost.
			close(court)
			return
		}
		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// Hit the ball back to the opposing player.
		court <- ball
	}
}
