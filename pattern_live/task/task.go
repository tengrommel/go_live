package main

import "sync"

type Worker interface {
	Work()
}

type Task struct {
	work chan Worker
	wg sync.WaitGroup
}

func New(maxGoroutines int) *Task {
	t := Task{work:make(chan Worker)}
	t.wg.Add(maxGoroutines)

	for i:=0;i<maxGoroutines;i++{
		go func() {
			for w := range t.work{
				w.Work()
			}
		}()
	}
	return &t
}

// Shutdown waits for all the goroutines to shutdown.
func (t *Task)Shutdown()  {
	close(t.work)
	t.wg.Wait()
}

func main() {
}
