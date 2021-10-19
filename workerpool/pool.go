package workerpool

import (
	"log"
	"sync"
)

type WorkFunc interface {
	Run() bool
}

type work struct {
	fn WorkFunc
}

type GoroutinePool struct {
	workers, ipw, maxItems int
	wg                     sync.WaitGroup
	queue                  chan work
	finished               chan bool
}

// Returns a worker pool with the given arguments
func NewGoroutinePool(workers, ipw, maxItems int) *GoroutinePool {
	gp := &GoroutinePool{
		workers:  workers,
		ipw:      ipw,
		maxItems: maxItems,
		queue:    make(chan work),
		finished: make(chan bool),
	}

	gp.AddWorkers(workers)
	return gp
}

// Waits for the workers to have finished their work.
func (gp *GoroutinePool) Close() {
	gp.wg.Wait()
}

// Adds a work to the worker pool queue.
func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn}
}

// Sends a signal to all workers to let them know all tasks have been scheduled.
func (gp *GoroutinePool) SetFinished() {
	for i := 0; i < gp.workers; i++ {
		go func() {
			gp.finished <- true
		}()
	}
}

// Adds the given amount of workers to the worker pool.
func (gp *GoroutinePool) AddWorkers(workers int) {
	gp.wg.Add(workers)
	mutex := &sync.Mutex{}
	totalCount := 0
	for i := 0; i < workers; i++ {
		go func(workerId int) {
			workerCount := 0
			processed := 0

		loop:
			for {
				mutex.Lock()
				if !(workerCount < gp.ipw && totalCount < gp.maxItems) {
					mutex.Unlock()
					break
				}

				select {
				case <-gp.finished:
					mutex.Unlock()
					break loop

				case job := <-gp.queue:
					added := job.fn.Run()
					if added {
						workerCount++
						totalCount++
					}

					processed++
				}

				mutex.Unlock()
			}
			log.Printf("Worker %d received %d items and returned %d items\n", workerId, processed, workerCount)
			gp.wg.Done()
		}(i + 1)
	}
}
