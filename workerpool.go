package main

import (
	"log"
)

// IWorkerPool is a contract for Worker Pool implementation
type IWorkerPool interface {
	Run()
	Addjob(job func())
}

// WorkerPool is a class with available workers and queued job information
type WorkerPool struct {
	maxWorker int
	queuedJob chan func()
}

// NewWorkerPool will create an instance of WorkerPool.
func NewWorkerPool(maxWorker int) *WorkerPool {
	return &WorkerPool{
		maxWorker: maxWorker,
		queuedJob: make(chan func()),
	}
}

func (wp *WorkerPool) Run() {
	wp.run()
}

func (wp *WorkerPool) AddJob(job func()) {
	wp.queuedJob <- job
}

func (wp *WorkerPool) GetQueuedjobs() int {
	return len(wp.queuedJob)
}

func (wp *WorkerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		workerId := i + 1
		log.Printf("worker %d has been spawned", workerId)

		go func(workerId int) {
			for job := range wp.queuedJob {
				log.Printf("worker %d start processing job", workerId)
				job()
				log.Printf("worker %d finish processing job", workerId)
			}
		}(workerId)
	}
}
