package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded    = errors.New("errors limit exceeded")
	ErrInvalidGoroutinesCount = errors.New("goroutines count must be greater than zero")
	ErrInvalidErrorsCount     = errors.New("errors count must be greater than or equal to zero")
	countErrors               int32
)

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidGoroutinesCount
	}
	if m <= 0 {
		return ErrInvalidErrorsCount
	}
	countErrors = 0
	taskCh := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(<-chan Task, int32) {
			defer wg.Done()
			consume(taskCh, int32(m))
		}(taskCh, int32(m))
	}
	wg.Wait()

	if countErrors >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func consume(jobs <-chan Task, maxCountErr int32) {
	for task := range jobs {
		if atomic.LoadInt32(&countErrors) >= maxCountErr {
			return
		}

		err := task()
		if err == nil {
			continue
		}

		atomic.AddInt32(&countErrors, 1)
		if atomic.LoadInt32(&countErrors) >= maxCountErr {
			return
		}
	}
}
