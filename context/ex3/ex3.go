package main

import (
	"time"
	"context"
	"fmt"
)

func main() {
	// Set a deadline.
	dealine := time.Now().Add(150 * time.Millisecond)

	// Create a context that is both manually cancellable and will signal a cancel at the specificed data/time.
	ctx, cancel := context.WithDeadline(context.Background(), dealine)
	defer cancel()

	// Create a channel to received a signal that work is done.
	ch := make(chan struct{})

	// Ask the goroutine to do some work for us.
	go func() {
			// Simulate work.
			time.Sleep(50*time.Millisecond)
			// Report the work is done.
			ch <- struct{}{}
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case <- ch:
		fmt.Println("work complete")
	case <- ctx.Done():
		fmt.Println("work cancelled")
	}
}
