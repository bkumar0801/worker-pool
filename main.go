package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	// For monitoring purpose.
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	// Start Worker Pool.
	totalWorker := 5
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 100
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1

		job := func() {
			time.Sleep(5 * time.Second)
			resultC <- result{id, id * 2}
		}

		log.Printf("created a job %d", id)

		wp.AddJob(job)
	}

	for r := range resultC {
		log.Printf("job %d has been finished with result %d", r.id, r.value)
	}

	<-waitC
}
