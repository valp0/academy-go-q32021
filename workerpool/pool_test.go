package workerpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type testTask struct {
	Name          string
	TaskProcessor func(...interface{})
}

func (t *testTask) Run() bool {
	t.TaskProcessor(t.Name)
	return true
}

func TestDispatcher(t *testing.T) {
	pool := NewGoroutinePool(7, 500, 1000)
	taskSize := 50
	taskCounter := 0

	wg := &sync.WaitGroup{}
	wg.Add(taskSize)

	//specific task
	sampleStringTaskFn := func(dm ...interface{}) {
		if input, ok := dm[0].(string); ok {
			time.Sleep(time.Microsecond)
			if input != "" {
				fmt.Printf("Finished %s\n", input)
			}
			taskCounter++
			wg.Done()
		}
	}

	var tasks []*testTask
	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, &testTask{
			Name:          fmt.Sprintf("task %v", v),
			TaskProcessor: sampleStringTaskFn,
		})
	}

	for _, task := range tasks {
		pool.ScheduleWork(task)
	}
	pool.Close()

	wg.Wait()

	if pool == nil {
		t.Fatal("pool is nil")
	}

	if taskCounter != taskSize {
		t.Fatal("not all tasks were performed")
	}
}
