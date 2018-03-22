# Go's Concurrency Building Blocks

## Goroutines 
> Goroutines are one of the most basic units of organization in a Go program, so it's important we understand what they are and how they work.

1. In fact, every Go program has at least one goroutine: the main goroutine, which is automatically created and started when the process begins.
2. In almost any program you'll probably find yourself reaching for a goroutine sooner or later to assist in solving your problems.


*What makes goroutines unique to Go are their deep integration with Go’s runtime.*

1. Goroutines don't define their own suspension or reentry points.
2. Go's runtime observes the runtime behavior of goroutines and automatically suspends
 them when they block and then resumes them when they become blocked.
3. In a way this makes them preemptable, but only at points where the goroutine has become blocked.

Coroutines, and thus goroutines, are implicitly concurrent constructs, but concurrency is not a property of a coroutine: 

**something must host several coroutines simultaneously and give each an opportunity to execute — otherwise, 
they wouldn’t be concurrent!**

Go follows a model of concurrency called the fork-join model.

1. The word fork refers to the fact that at any point in the program, it can split off a child branch of 
   execution to be run concurrently with its parent. 

2. The word join refers to the fact that at some point in the future, these concurrent branches of 
   execution will join back together. 

