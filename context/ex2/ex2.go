package main

import (
	"context"
	"time"
	"fmt"
)

func main() {
	// Create a context that is cancellable only manually.
	ctx, cancel := context.WithCancel(context.Background())

	// The cancel function must be called regardless of the outcome.
	defer cancel()

	go func() {
			// Simulate work.
			time.Sleep(150*time.Millisecond)
			cancel()
	}()

	select {
	case <- time.After(100*time.Millisecond):
		fmt.Println("moving on")
	case <- ctx.Done():
		fmt.Println("work complete")
	}

}
