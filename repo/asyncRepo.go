package repo

import (
	"encoding/csv"
	"io"
	"os"
	"sync"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/workerpool"
)

type asyncRepo struct{}

// Returns a local repository to handle async database operations.
func NewAsyncRepo() asyncRepo {
	return asyncRepo{}
}

// Returns elements matching type, items and items_per_worker.
func (r asyncRepo) Filter(t, items, ipw, path string) ([]common.Element, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []common.Element{}, err
	}
	defer f.Close()

	filtered := []common.Element{}

	parity, err := checkType(t)
	if err != nil {
		return filtered, err
	}

	itemAmount, err := checkItems(items)
	if err != nil {
		return filtered, err
	}

	itemsPerWorker, err := checkIpw(ipw)
	if err != nil {
		return filtered, err
	}

	// We create a worker pool with n workers.
	n := getMaxWorkers()
	pool := workerpool.NewGoroutinePool(n, itemsPerWorker, itemAmount)
	mutex := &sync.Mutex{}

	filterProcessor := func(element common.Element, parity string) bool {
		mutex.Lock()
		added := false
		switch parity {
		case "even":
			if element.Key%2 == 0 {
				filtered = append(filtered, element)
				added = true
			}
		case "odd":
			if element.Key%2 != 0 {
				filtered = append(filtered, element)
				added = true
			}
		default:
			filtered = append(filtered, element)
			added = true
		}
		mutex.Unlock()
		return added
	}

	var errStatus error = nil
	csvReader := csv.NewReader(f)
	go func() {
		i := 0
		for {
			i++
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				errStatus = err
				break
			}

			element, err := getElement(record, i)
			if err != nil {
				errStatus = err
				break
			}

			task := getTask(element, filterProcessor, parity)
			pool.ScheduleWork(task)
		}

		pool.SetFinished()
	}()

	pool.Close()
	if errStatus != nil {
		return []common.Element{}, errStatus
	}

	sortElements(filtered)
	return filtered, nil
}

type filterTask struct {
	parity  string
	element common.Element
	filter  func(common.Element, string) bool
}

// Filters an element of a filterTask using its parity (even/odd).
func (ft *filterTask) Run() bool {
	return ft.filter(ft.element, ft.parity)
}
