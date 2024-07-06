package pipeline

import (
	"context"
	"fmt"
	"sync"
)

type TaskFunc func(ctx context.Context, index int) (any, error)

type Result struct {
	Index int
	Value any
	Err   error
}

func Parallelize(ctx context.Context, numTasks int, task TaskFunc) ([]Result, error) {
	results := make([]Result, numTasks)
	resultCh := make(chan Result, numTasks)
	errCh := make(chan error, 1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go func(index int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					select {
					case errCh <- fmt.Errorf("panic in task %d: %v", index, r):
					default:
					}
				}
			}()

			value, err := task(ctx, index)
			select {
			case resultCh <- Result{Index: index, Value: value, Err: err}:
			case <-ctx.Done():
			}

			if err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultCh)
		close(errCh)
	}()

	for i := 0; i < numTasks; i++ {
		select {
		case res := <-resultCh:
			results[res.Index] = res
		case err := <-errCh:
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return results, nil
}
