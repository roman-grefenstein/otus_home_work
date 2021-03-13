package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded    = errors.New("errors limit exceeded")
	errInvalidGoroutinesCount = errors.New("goroutines count must be greater than zero")
	countErrors               int32
)

type Task func() error

// Run starts tasks in `n` goroutines and stops its work when receiving `m` errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return errInvalidGoroutinesCount
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
		go func(int32) {
			defer wg.Done()
			consume(taskCh, int32(m))
		}(int32(m))
	}
	wg.Wait()

	if maxErrorReached(int32(m)) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func consume(jobs <-chan Task, maxCountErr int32) {
	for task := range jobs {
		if maxErrorReached(maxCountErr) {
			return
		}

		err := task()
		if err == nil {
			continue
		}

		atomic.AddInt32(&countErrors, 1)
		if maxErrorReached(maxCountErr) {
			return
		}
	}
}

func maxErrorReached(maxCountErr int32) bool {
	return maxCountErr > 0 && atomic.LoadInt32(&countErrors) >= maxCountErr
}
