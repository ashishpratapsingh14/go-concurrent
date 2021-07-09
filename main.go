package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main() {
	JobQueue = make(chan Job, 10000)
	dispatcher := NewDispatcher(500)
	dispatcher.Run()
	doFunc := func(data interface{}) error {
		fmt.Println(data)
		return nil
	}
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(c int) {
			defer wg.Done()
			start := c * 100
			end := (c + 1) * 100
			for j := start; j < end; j++ {
				// let's create a job with the payload
				work := Job{Payload: j, Executor: doFunc}
				// Push the work onto the queue.
				JobQueue <- work
			}
		}(i)
	}
	wg.Wait()
}
