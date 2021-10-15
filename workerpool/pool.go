package workerpool

import (
	"log"
	"sync"
	"time"
)

type WorkFunc interface {
	Run() bool
}

type work struct {
	fn WorkFunc
}

type GoroutinePool struct {
	queue    chan work
	ipw      int
	maxItems int
	wg       sync.WaitGroup
}

// Returns a worker pool with the given arguments
func NewGoroutinePool(workers, ipw, maxItems int) *GoroutinePool {
	gp := &GoroutinePool{
		queue:    make(chan work),
		ipw:      ipw,
		maxItems: maxItems,
	}

	gp.AddWorkers(workers)
	return gp
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn}
}

func (gp *GoroutinePool) AddWorkers(workers int) {
	gp.wg.Add(workers)
	mutex := &sync.Mutex{}
	totalCount := 0
	for i := 0; i < workers; i++ {
		go func(workerId int) {
			itemCount := 0
			processed := 0

			for job := range gp.queue {
				mutex.Lock()
				if gp.ipw <= itemCount || gp.maxItems <= totalCount {
					mutex.Unlock()
					time.Sleep(time.Nanosecond)
					processed++
					continue
				}

				added := job.fn.Run()
				if added {
					itemCount++
					totalCount++
				}
				mutex.Unlock()

				processed++
			}
			log.Printf("Worker %d processed %d items and returned %d items\n", workerId, processed, itemCount)
			gp.wg.Done()
		}(i + 1)
	}
}
